1.After unzipping, place this folder parking_lot in your GO path directory, with directory structure as $GOPATH/src/github.com/parking_lot.

2.cd to $GOPATH/src/github.com/parking_lot and run 'bin/setup'.
   This will automatically install dependencies following which a 'vendor' folder will get created and perform a go build and then perform some functional checks.

3.Run bin/parking_lot parking_lot_file_inputs.txt to execute instructions from file.

4.Run bin/parking_lot to execute instructions from command line.