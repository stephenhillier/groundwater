pipeline {
  environment {}
  agent any
  stages {
    stage('Start pipeline') {
      steps {
        script {
          abortAllPreviousBuildInProgress(currentBuild)
        }
      }
    }

    // start a new build for this pull request.
    // 
    // todo: output unit tests to jenkins output. Currently they will fail the pipeline run
    // but without any output.
    stage('Test & build image') {
      steps {
        script {
          openshift.withCluster() {
            // create a new build config for WELLS service
            // unit tests are run during build (application will not be built if unit tests fail)
            if ( !openshift.selector("bc", "wells-builder").exists() ) {
              echo "Creating a new build config"
              openshift.newBuild("https://github.com/stephenhillier/groundwater.git", "--strategy=docker", "--name=wells-builder" "--context-dir=wells/")
            } else {
              echo "Starting build"
              openshift.selector("bc", "wells-builder").startBuild("--wait")
            }

            echo "Waiting for builds from buildconfig wells-builder to finish"
            def lastBuildNumber = openshift.selector("bc", "wells-builder").object().status.lastVersion
            def lastBuild = openshift.selector("build", "wells-builder-${lastBuildNumber}")
            timeout(10) {
              lastBuild.untilEach(1) {
                return it.object().status.phase == "Complete"
              }
            }

            // start building an application image.
            // this is a chained build; only the application binary will be brought forward from the builder image.
            // the image can only be used as an executable
            if ( !openshift.selector("bc", "wells").exists() ) {
              echo "Creating new application build config"
              openshift.newBuild("alpine:3.8", "--source-image=wells-builder", "--allow-missing-imagestream-tags", "--name=wells", "--source-image-path=/go/bin/wells:.", """--dockerfile='FROM alpine:3.8
              RUN mkdir -p /app
              COPY wells /app/wells
              ENTRYPOINT [\"/app/wells\"]
              '""")
            } else {
              echo "Creating application image"
              openshift.selector("bc", "wells").startBuild("--wait")
            }

            echo "Waiting for application build from wells to finish"
            def lastAppBuildNumber = openshift.selector("bc", "wells").object().status.lastVersion
            def lastAppBuild = openshift.selector("build", "wells-${lastAppBuildNumber}")
            timeout(10) {
              lastAppBuild.untilEach(1) {
                return it.object().status.phase == "Complete"
              }
            }

            // the dev deployment will automatically run as soon as a new image is tagged as `dev`
            echo "Successfully built image: tagging as new dev image"
            openshift.tag("wells:latest", "wells:dev")

          }
        }
      }
    }

    // Deployment to dev happens automatically when a new image is tagged `dev`.
    // This stage monitors the newest deployment for pods/containers to report back as ready.
    stage('Deploy to dev') {
      steps {
        script {
          openshift.withCluster() {

            // if a deployment config does not exist for this pull request, create one
            if ( !openshift.selector("dc", "wells-dev").exists() ) {
              echo "Creating a new deployment config for WELLS"
              openshift.newApp("wells:dev", "--name=wells-dev").narrow("dc").expose("--port=8000")
            }

            echo "Waiting for deployment to dev..."
            def newVersion = openshift.selector("dc", "wells-dev").object().status.latestVersion

            // find the pods for the newest deployment
            def pods = openshift.selector('pod', [deployment: "wells-dev-${newVersion}"])

            // wait until each container in this deployment's pod reports as ready
            timeout(10) {
              pods.untilEach(1) {
                return it.object().status.containerStatuses.every {
                  it.ready
                }
              }
              echo "Deployment successful!"
            }
          }
        }
      }
    }
  }
}

