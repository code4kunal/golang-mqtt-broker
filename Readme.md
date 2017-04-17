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
 - Currently, the test scenario allows a total of 4 readings in a single run.
  
## Execution
  
  Set mqtt broker up and running:
 ```sh
 $ docker run -d -p 1883:1883 -p 8883:8883 prologic/mosquitto

 ``` 
  Set up project
   ```sh
   $ git clone https://github.com/code4kunal/golang-task.git
   $ cd golang-task
   $ go build && ./golang-task 
  
   ``` 