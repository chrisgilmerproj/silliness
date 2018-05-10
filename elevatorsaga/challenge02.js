/*
 * Challenge #2: Transport 20 people in 60 seconds or less
 */

{
    init: function(elevators, floors) {

        elevators.forEach(function(elevator) {
            // When the elevator has completed all its tasks and is not doing anything.
            elevator.on("idle", function() {
                // Go the the main floor
                elevator.goToFloor(0);
            });

            // When a passenger has pressed a button inside the elevator.
            elevator.on("floor_button_pressed", function() {
                elevator.getPressedFloors().forEach(function(floorNum) {
                    elevator.goToFloor(floorNum);
                });
            });

            // Slightly before the elevator will pass a floor.
            elevator.on("passing_floor", function() {
                elevator.destinationQueue.sort();
                elevator.checkDestinationQueue();
            });

            // When the elevator has arrived at a floor.
            elevator.on("stopped_at_floor", function() {
            });
        });

        floors.forEach(function(floor) {
            // When someone has pressed the up button at a floor.
            floor.on("up_button_pressed", function() {
                elevators.forEach(function(elevator) {
                    if (elevator.destinationDirection == "up" ) {
                        elevator.goToFloor(floor.floorNum);
                    } else if (elevator.destinationDirection == "stopped") {
                        elevator.goToFloor(floor.floorNum);
                    };
                });
            });

            // When someone has pressed the down button at a floor.
            floor.on("down_button_pressed", function() {
                elevators.forEach(function(elevator) {
                    if (elevator.destinationDirection == "down" ) {
                        elevator.goToFloor(floor.floorNum);
                    };
                });
            });
        });

    },
    update: function(dt, elevators, floors) {
        // We normally don't need to do anything here
    }
}
