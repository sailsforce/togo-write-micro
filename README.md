[![GitHub Release](https://img.shields.io/github/release/sailsforce/togo-write-micro.svg?style=flat)]() [![CodeFactor](https://www.codefactor.io/repository/github/sailsforce/togo-write-micro/badge)](https://www.codefactor.io/repository/github/sailsforce/togo-write-micro) ![security rating](./badges/securityRating.svg) ![vulnerabilities](./badges/vulnerabilities.svg) [![codecov](https://codecov.io/gh/sailsforce/inv-write-micro/branch/main/graph/badge.svg?token=U1Q38I84A2)](https://codecov.io/gh/sailsforce/inv-write-micro)

# togo-write-micro
Togo write micro-service

## Install pre-commit
```
brew install pre-commit
```

## Run Application Using Doppler
```
go build -o bin/togo-write -v .
doppler run --command="./bin/togo-write"
```

## Run Tests Using Doppler
```
doppler run --command="go test -v ./tests -coverprofile=./coverage.out -coverpkg ./..."
```

## See Coverage Report
```
go tool cover -html=./tests/coverage.out
```

## Upload to Codecov
``` 
./codecov -t ${CODECOV_TOKEN}
``` 

## Generate Coverage and Test Reports for SonarQube
```
go test ./tests -coverprofile=./tests/coverage.out -coverpkg ./... -json > ./tests/test-report.out
```

## Run SonarQube
```
brew install sonar-scanner

sonar-scanner \
  -Dsonar.projectKey=togo-write-micro \
  -Dsonar.sources=. \
  -Dsonar.host.url=${SONAR_HOST_URL} \
  -Dsonar.login=${SONAR_LOGIN}
```

## Update SonarQube badges
```
curl ${badge_url} > ./badges/securityRating.svg
curl ${badge_url} > ./badges/vulnerabilities.svg
```
