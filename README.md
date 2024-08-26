# Lemin Project

## Description

The Lemin Project simulates the movement of ants through a maze to evaluate and optimize different pathfinding strategies. The goal is to minimize the number of turns taken by the ants to reach the end room.

## Project Logic

The Lemin Project is designed to simulate and optimize the movement of ants through a maze. Here's an overview of the logic and key components used in the project:

### 1. **Parsing the Input**

The project begins by parsing an input file to extract essential information about the maze. This includes:
- **Room Coordinates**: The positions of each room in the maze.
- **Distance**: Calculte distance according to room coordinates
- **Tunnels**: Connections between rooms that define possible paths for the ants.

### 2. **Pathfinding Algorithms**

Two primary strategies are used to determine the ants' movements:
- **Path Sorting by Number of Rooms**: Paths are sorted based on the number of rooms they contain. Shorter paths (with fewer rooms) are prioritized.
- **Path Sorting by Distance**: Paths are sorted based on the distance between rooms. Shorter distances are preferred to minimize travel time.

### 3. **Simulating Ant Movements**

Ant movements are simulated based on the selected pathfinding strategy:
- **Movement Decision**: For each ant, the next room is selected based on the current strategy (number of rooms or distance).
- **Handling Congestion**: 
  - **Tunnel Congestion**: A map tracks occupied tunnels to prevent multiple ants from using the same tunnel simultaneously. This ensures that ants do not block each other’s progress.
  - **Room Occupancy**: A map also tracks which rooms are occupied. Each room’s occupancy status is updated whenever an ant moves in or out. This ensures that no two ants occupy the same room at the same time and helps in managing the movement flow effectively.
- **Ant Position Updates**: The position of each ant is updated, and the number of ants in the start and end rooms is adjusted accordingly.

### 4. **Comparing Strategies**

The project compares the performance of different pathfinding strategies by:
- **Running Simulations**: The maze is simulated using both sorting criteria (number of rooms and distance).
- **Tracking Turns**: Only the number of turns required for all ants to reach the end room is tracked and compared.
- **Output Results**: The results are printed, showing the movements of the ants and the number of turns required for each strategy.

### 5. **Displaying Results**

The results of the simulation are displayed in a readable format:
- **Ant Movements**: The sequence of movements for each ant is shown per turn.
- **Turn Count**: The total number of turns required for all ants to reach the end room is displayed.

This approach ensures that the system not only finds efficient paths but also evaluates and compares different strategies to determine the optimal solution for the given maze.

## Installation

To get started with the Lemin Project, follow these steps:

1. **Clone the Repository**

   ```bash
   git clone https://zone01normandie.org/git/ejean/lem-in.git
   ```
2. **Navigate to the project Directory**
    ```bash
    cd lemin
    ```
3. **Run program**
    ```bash
    go run main.go exemple/exemple00.txt
    ```
- You can try with different examples present in the "exemples/" directory.
