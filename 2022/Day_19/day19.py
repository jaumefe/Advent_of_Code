with open("Input/test.txt") as f:
    lines = f.readlines()

for i in range(len(lines)):
    blueprint = lines[i].strip().replace(':','').split()
    iD = int(blueprint[1])
    oreRobotCost = int(blueprint[6])
    clayRobotCost = int(blueprint[12])
    obsRobotCost = [int(blueprint[18]), int(blueprint[21])]
    geodeRobotCost = [int(blueprint[27]), int(blueprint[30])]
    oreRobot = 1
    clayRobot = 0
    obsRobot = 0
    geodeRobot = 0
    minutes = 0
    ore = 0
    clay = 0
    obs = 0
    geode = 0

    while minutes < 24:
        clayRobot_temp = 0
        obsRobot_temp = 0
        geodeRobot_temp = 0
        ore_temp = 0
        clay_temp = 0
        obs_temp = 0
        geode_temp = 0
        ore_temp += oreRobot
        clay_temp += clayRobot
        obs_temp += obsRobot
        geode_temp += geodeRobot
        minutes += 1
        if ore >= geodeRobotCost[0] and obs >= geodeRobotCost[1]:
            if (ore / geodeRobotCost[0]) <= (obs / geodeRobotCost[1]):
                geodeRobot_temp += int(ore / geodeRobotCost[0])
            else:
                geodeRobot_temp += int(obs / geodeRobotCost[1])
            ore -= geodeRobot_temp * geodeRobotCost[0]
            obs -= geodeRobot_temp * geodeRobotCost[1]
        elif ore >= obsRobotCost[0] and clay >= obsRobotCost[1]:
            if (ore / obsRobotCost[0]) <= (clay / obsRobotCost[1]):
                obsRobot_temp += int(ore / obsRobotCost[0])
            else:
                obsRobot_temp += int(clay / obsRobotCost[1])
            ore -= obsRobot_temp * obsRobotCost[0]
            clay -= obsRobot_temp * obsRobotCost[1]
        elif ore >= clayRobotCost:
            clayRobot_temp += int(ore/clayRobotCost)
            ore -= clayRobot_temp * clayRobotCost
        ore += ore_temp
        clay += clay_temp
        obs += obs_temp
        geode += geode_temp
        clayRobot += clayRobot_temp
        obsRobot += obsRobot_temp
        geodeRobot += geodeRobot_temp
    print(geode*iD)


