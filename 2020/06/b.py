#!/usr/bin/python3

import sys

def main():
    all_input = sys.stdin.read()
    groups = all_input.split('\n\n')

    count = 0

    for group in groups:
        accumulator = None
        for line in group.split('\n'):
            line = line.strip()
            if len(line) == 0:
                break
            if accumulator is None:
                accumulator = set(line)
            else:
                accumulator.intersection_update(set(line))
        count += len(accumulator)

    print(count)

if __name__ == '__main__':
    sys.exit(main())
