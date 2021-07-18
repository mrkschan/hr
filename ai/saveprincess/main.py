#!/usr/bin/python

def displayPathtoPrincess(n,grid):
    moves = (n-1)//2
    if grid[0][0] == 'p':
        print("UP\nLEFT\n" * moves)
    elif grid[0][n-1] == 'p':
        print("UP\nRIGHT\n" * moves)
    elif grid[n-1][0] == 'p':
        print("DOWN\nLEFT\n" * moves)
    elif grid[n-1][n-1] == 'p':
        print("DOWN\nRIGHT\n" * moves)


m = int(input())
grid = []
for i in range(0, m):
    grid.append(input().strip())

displayPathtoPrincess(m,grid)
