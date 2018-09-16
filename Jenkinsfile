pipeline {
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
              openshift.newBuild("https://github.com/stephenhillier/groundwater.git", "--strategy=docker", "--name=wells-builder", "--context-dir=wells/", "-l app-name=groundwater")
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
              openshift.newBuild("openshift/alpine:3.7", "--source-image=wells-builder", "--allow-missing-imagestream-tags", "--name=wells", "--source-image-path=/go/bin/wells:.", "-l app-name=groundwater", """--dockerfile='FROM openshift/alpine:3.7
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

            // Aquifers

            // create a new build config for aquifers service
            // unit tests are run during build (application will not be built if unit tests fail)
            if ( !openshift.selector("bc", "aquifers-builder").exists() ) {
              echo "Creating a new build config"
              openshift.newBuild("https://github.com/stephenhillier/groundwater.git", "--strategy=docker", "--name=aquifers-builder", "--context-dir=aquifers/", "-l app-name=groundwater")
            } else {
              echo "Starting build"
              openshift.selector("bc", "aquifers-builder").startBuild("--wait")
            }

            echo "Waiting for builds from buildconfig aquifers-builder to finish"
            def lastBuildNumber = openshift.selector("bc", "aquifers-builder").object().status.lastVersion
            def lastBuild = openshift.selector("build", "aquifers-builder-${lastBuildNumber}")
            timeout(10) {
              lastBuild.untilEach(1) {
                return it.object().status.phase == "Complete"
              }
            }

            // start building an application image.
            // this is a chained build; only the application binary will be brought forward from the builder image.
            // the image can only be used as an executable
            if ( !openshift.selector("bc", "aquifers").exists() ) {
              echo "Creating new application build config"
              openshift.newBuild("openshift/alpine:3.7", "--source-image=aquifers-builder", "--allow-missing-imagestream-tags", "--name=aquifers", "--source-image-path=/go/bin/aquifers:.", "-l app-name=groundwater", """--dockerfile='FROM openshift/alpine:3.7
              RUN mkdir -p /app
              COPY aquifers /app/aquifers
              ENTRYPOINT [\"/app/aquifers\"]
              '""")
            } else {
              echo "Creating application image"
              openshift.selector("bc", "aquifers").startBuild("--wait")
            }

            echo "Waiting for application build from aquifers to finish"
            def lastAppBuildNumber = openshift.selector("bc", "aquifers").object().status.lastVersion
            def lastAppBuild = openshift.selector("build", "aquifers-${lastAppBuildNumber}")
            timeout(10) {
              lastAppBuild.untilEach(1) {
                return it.object().status.phase == "Complete"
              }
            }


            // the dev deployment will automatically run as soon as a new image is tagged as `dev`
            echo "Successfully built image: tagging as new dev image"
            openshift.tag("aquifers:latest", "aquifers:dev")

            // API

            // create a new build config for aquifers service
            // unit tests are run during build (application will not be built if unit tests fail)
            if ( !openshift.selector("bc", "api-builder").exists() ) {
              echo "Creating a new build config"
              openshift.newBuild("https://github.com/stephenhillier/groundwater.git", "--strategy=docker", "--name=api-builder", "--context-dir=api/", "-l app-name=groundwater")
            } else {
              echo "Starting build"
              openshift.selector("bc", "api-builder").startBuild("--wait")
            }

            echo "Waiting for builds from buildconfig api-builder to finish"
            def lastBuildNumber = openshift.selector("bc", "api-builder").object().status.lastVersion
            def lastBuild = openshift.selector("build", "api-builder-${lastBuildNumber}")
            timeout(10) {
              lastBuild.untilEach(1) {
                return it.object().status.phase == "Complete"
              }
            }

            // start building an application image.
            // this is a chained build; only the application binary will be brought forward from the builder image.
            // the image can only be used as an executable
            if ( !openshift.selector("bc", "api").exists() ) {
              echo "Creating new application build config"
              openshift.newBuild("openshift/alpine:3.7", "--source-image=api-builder", "--allow-missing-imagestream-tags", "--name=api", "--source-image-path=/go/bin/api:.", "-l app-name=groundwater", """--dockerfile='FROM openshift/alpine:3.7
              RUN mkdir -p /app
              COPY api /app/api
              ENTRYPOINT [\"/app/api\"]
              '""")
            } else {
              echo "Creating application image"
              openshift.selector("bc", "api").startBuild("--wait")
            }

            echo "Waiting for application build from api to finish"
            def lastAppBuildNumber = openshift.selector("bc", "api").object().status.lastVersion
            def lastAppBuild = openshift.selector("build", "api-${lastAppBuildNumber}")
            timeout(10) {
              lastAppBuild.untilEach(1) {
                return it.object().status.phase == "Complete"
              }
            }


            // the dev deployment will automatically run as soon as a new image is tagged as `dev`
            echo "Successfully built image: tagging as new dev image"
            openshift.tag("api:latest", "api:dev")

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
            if ( !openshift.selector("dc", "wells").exists() ) {
              echo "Creating a new deployment config for WELLS"
              openshift.newApp("wells:dev aquifers:dev api:dev", "-l app-name=groundwater").narrow("dc").expose("--port=8000")
            }

            echo "Waiting for deployment to dev..."
            def newVersion = openshift.selector("dc", "wells").object().status.latestVersion

            // find the pods for the newest deployment
            def pods = openshift.selector('pod', [deployment: "wells-${newVersion}"])

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

