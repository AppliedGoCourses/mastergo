#!/usr/bin/env sh
ROWS=1000
COLS=10
DELAY=1ms
echo "Benchmarking serial code (rows=$ROWS, cols=$COLS, delay=$DELAY)"
go test -run=NONE -bench=. -args -rows=$ROWS -cols=$COLS -delay=$DELAY >bench.out
cd solution1
echo Benchmarking solution 1
go test -run=NONE -bench=. -args -rows=$ROWS -cols=$COLS -delay=$DELAY >bench.out
cd ../solution2
echo Benchmarking solution 2
go test -run=NONE -bench=. -args -rows=$ROWS -cols=$COLS -delay=$DELAY >bench.out
cd ..
echo Changes from serial to solution 1
benchcmp bench.out solution1/bench.out
echo Changes from serial to solution 2
benchcmp bench.out solution2/bench.out
echo Changes from solution 1 to solution 2
benchcmp solution1/bench.out solution2/bench.out