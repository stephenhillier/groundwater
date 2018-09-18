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

            echo "Starting builds"
            openshift.selector("bc", "wells-builder").startBuild()
            openshift.selector("bc", "aquifers-builder").startBuild()
            openshift.selector("bc", "api-builder").startBuild()


            echo "Waiting for builds from buildconfig wells-builder to finish"
            def lastWellBuildNumber = openshift.selector("bc", "wells-builder").object().status.lastVersion
            def lastWellBuild = openshift.selector("build", "wells-builder-${lastWellBuildNumber}")
            timeout(10) {
              lastWellBuild.untilEach(1) {
                return it.object().status.phase == "Complete"
              }
            }

            echo "Waiting for builds from buildconfig aquifers-builder to finish"
            def lastAquiferBuildNumber = openshift.selector("bc", "aquifers-builder").object().status.lastVersion
            def lastAquiferBuild = openshift.selector("build", "aquifers-builder-${lastAquiferBuildNumber}")
            timeout(10) {
              lastAquiferBuild.untilEach(1) {
                return it.object().status.phase == "Complete"
              }
            }

            echo "Waiting for builds from buildconfig api-builder to finish"
            def lastAPIBuildNumber = openshift.selector("bc", "api-builder").object().status.lastVersion
            def lastAPIBuild = openshift.selector("build", "api-builder-${lastAPIBuildNumber}")
            timeout(10) {
              lastAPIBuild.untilEach(1) {
                return it.object().status.phase == "Complete"
              }
            }

            // start building an application image.
            // this is a chained build; only the application binary will be brought forward from the builder image.

            echo "Creating application images"
            openshift.selector("bc", "wells").startBuild("--wait")
            openshift.selector("bc", "aquifers").startBuild("--wait")
            openshift.selector("bc", "api").startBuild("--wait")


            echo "Waiting for application build from wells to finish"
            def lastWellAppNumber = openshift.selector("bc", "wells").object().status.lastVersion
            def lastWellApp = openshift.selector("build", "wells-${lastWellAppNumber}")
            timeout(10) {
              lastWellApp.untilEach(1) {
                return it.object().status.phase == "Complete"
              }
            }

            echo "Waiting for application build from aquifers to finish"
            def lastAquiferAppBuildNumber = openshift.selector("bc", "aquifers").object().status.lastVersion
            def lastAquiferAppBuild = openshift.selector("build", "aquifers-${lastAquiferAppBuildNumber}")
            timeout(10) {
              lastAquiferAppBuild.untilEach(1) {
                return it.object().status.phase == "Complete"
              }
            }

            echo "Waiting for application build from api to finish"
            def lastAPIAppBuildNumber = openshift.selector("bc", "api").object().status.lastVersion
            def lastAPIAppBuild = openshift.selector("build", "api-${lastAPIAppBuildNumber}")
            timeout(10) {
              lastAPIAppBuild.untilEach(1) {
                return it.object().status.phase == "Complete"
              }
            }


            // the dev deployment will automatically run as soon as a new image is tagged as `dev`
            echo "Successfully built images: tagging as new dev image"



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

            openshift.tag("api:latest", "api:dev")
            openshift.tag("aquifers:latest", "aquifers:dev")
            openshift.tag("wells:latest", "wells:dev")

            sleep(5)

            def newVersion = openshift.selector("dc", "api").object().status.latestVersion

            // find the pods for the newest deployment
            def pods = openshift.selector('pod', [deployment: "api-${newVersion}"])

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

