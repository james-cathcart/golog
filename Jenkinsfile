node('workers') {

    stage('Checkout') {
        scmInfo = checkout scm
        env.GIT_COMMIT = scmInfo.GIT_COMMIT
    }

    def testImage = docker.build("golog-test", "-f test.Dockerfile .")

    stage('Test') {
        testImage.inside {
            sh 'go test ./...'
        }
    }
}