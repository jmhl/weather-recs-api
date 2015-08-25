from math import sqrt
from recommendations import critics

def sim_distance(prefs, p1, p2):
    '''
    Calculates Euclidian distance score
    '''

    si = {}
    for item in prefs[p1]:
        if item in prefs[p2]:
            si[item] = 1

    if len(si) == 0:
        return 0

    sum_of_squares = 0
    for item in prefs[p1]:
        if item in prefs[p2]:
            prefs_p1, prefs_p2 = prefs[p1][item], prefs[p2][item]
            sum_of_squares += pow(prefs_p1 - prefs_p2, 2)

    return 1 / (1 + sum_of_squares)

print(sim_distance(critics, 'Lisa Rose','Gene Seymour'))
