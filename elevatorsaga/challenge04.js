/*
 * Challenge #4: Transport 28 people in 60 seconds or less
 */
{
    init: function(elevators, floors) {

        const waitingFloorsUp = [];
        const waitingFloorsDown = [];

        elevators.forEach(function(elevator) {
            // When the elevator has completed all its tasks and is not doing anything.
            elevator.on("idle", function() {
                // If people are waiting then go to the top floor, else go home
                if (waitingFloorsUp.length > 0) {
                    console.log(waitingFloorsUp);
                    elevator.goToFloor(Math.max(...waitingFloorsUp));
                } else if (waitingFloorsDown.length > 0) {
                    console.log(waitingFloorsDown);
                    elevator.goToFloor(Math.max(...waitingFloorsDown));
                } else {
                    // Go the the main floor
                    elevator.goToFloor(0);
                }
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
                const fNum = floor.floorNum();
                if (waitingFloorsUp.indexOf(fNum)) {
                    waitingFloorsUp.push(fNum);
                }
            });

            // When someone has pressed the down button at a floor.
            floor.on("down_button_pressed", function() {
                const fNum = floor.floorNum();
                if (waitingFloorsDown.indexOf(fNum)) {
                    waitingFloorsDown.push(fNum);
                }
            });
        });

    },
    update: function(dt, elevators, floors) {
        // We normally don't need to do anything here
    }
}
