# Wattx-Task


## Problem Statement

Emulate a heating service in which sensors in a room records temperature and periodically publish readings in a specific format.
After receiving the values, openness of the valve is determined and again published.

## Solution
 
 A message broker is used to publish and subscribe messages in the IOT environment. 
 Some assumptions are made to emulate the real world scenario.
 
 - A room consists of 10 sensors.
 - Every sensors randomly publishes the recorded reading attached to a given topic.
 - Every 20 seconds(periodically) a reading is published.
 - Currently, the test scenario allows a total of 4 readings ina  single run.
  
## Execution
  
  Set mqtt broker up and running:
 ```sh
 $ docker run -d -p 1883:1883 -p 8883:8883 prologic/mosquitto

 ``` 
 
```
Give examples
```

### Installing

A step by step series of examples that tell you have to get a development env running

Say what the step will be

```
Give the example
```

And repeat

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* [Dropwizard](http://www.dropwizard.io/1.0.2/docs/) - The web framework used
* [Maven](https://maven.apache.org/) - Dependency Management
* [ROME](https://rometools.github.io/rome/) - Used to generate RSS Feeds

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Billie Thompson** - *Initial work* - [PurpleBooth](https://github.com/PurpleBooth)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone who's code was used
* Inspiration
* etc
