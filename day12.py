rows = open("day11/input.txt").read().split("\n")
[x for x in range(10) if x % 2 == 0]

map = [left.split(" ")[0] for left in rows]
hotSpots = [sorted([int(item) for item in (left.split(" ")[1]).split(",")], reverse=True) for left in rows]
print(map)
print(hotSpots)

total = 0

# Part 1
def Solver(map, hotSpots):
   
    hotSpotSize = hotSpots.pop(0)

    for mapIndex in range(len(map) - hotSpotSize):
        # Check if there can be a hotspot here
        print("Checking function map: {}, hotSpotString: {}, hotSpotSize: {}, mapIndex: {}, rem HS: {}".format(map, map[mapIndex:mapIndex + hotSpotSize], hotSpotSize, mapIndex, hotSpots))
        if checkHotSpot(map, hotSpotSize, mapIndex):
            # Possible, so need to make new string consider there is a hotspot there
            # Check if hotspot is at end of string
            print("Found hotspot at index: {} with size {} in {}".format(mapIndex, hotSpotSize, map))
            if hotSpotSize + mapIndex + 1 == len(map):
                print("Found HS at end")
                if len(hotSpots) == 0:
                    print("No more hotspots to solve for")
                    return 1
                else:
                    newMap = map[:mapIndex]
                    Solver(newMap, hotSpots.copy())
            # Check if hotspot is at beginning or middle of string and has a . or ? after it
            elif map[hotSpotSize + mapIndex] in ".?":
                LeftMap = map[:mapIndex]
                RightMap = map[hotSpotSize + mapIndex + 1:] # +1 to skip the . or ?
                print("Found HS in beg/mid. Left: {}, Right: {}".format(LeftMap, RightMap))
                if len(hotSpots) == 0:
                    print("No more hotspots to solve for")
                    return 1
                else:
                    Solver(LeftMap, hotSpots.copy())
                    Solver(RightMap, hotSpots.copy())
            
    return 0


def checkHotSpot(map, hotSpotSize, mapIndex):
    if all(char in ['#', '?'] for char in map[(mapIndex+1):mapIndex + hotSpotSize]):
        return True
    else:
        return False

# ???.### 1,1,3
# .??..??...?##. 1,1,3
# ?#?#?#?#?#?#?#? 1,3,1,6
# ????.#...#... 4,1,1
# ????.######..#####. 1,6,5
# ?###???????? 3,2,1
ans = Solver("???.###", [3,1,1])
print(ans)