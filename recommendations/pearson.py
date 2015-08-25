from math import sqrt
from recommendations import critics

def sim_pearson(prefs, p1, p2):
    '''
    Calculates the Pearson coefficient
    '''

    si = {}
    for item in prefs[p1]:
        if item in prefs[p2]:
            si[item] = 1

    n = len(si)

    if n == 0:
        return 0

    sum1 = sum([prefs[p1][item] for item in si])
    sum2 = sum([prefs[p2][item] for item in si])

    sum1Sq = sum([pow(prefs[p1][item], 2) for item in si])
    sum2Sq = sum([pow(prefs[p2][item], 2) for item in si])

    sum_of_products = sum([prefs[p1][item] * prefs[p2][item] for item in si])

    num = sum_of_products - (sum1 * sum2 / n)
    den = sqrt((sum1Sq - pow(sum1, 2) / n) * (sum2Sq - pow(sum2, 2) / n))

    if den == 0:
        return 0

    coefficient = num / den

    return coefficient

print(sim_pearson(critics, 'Lisa Rose','Gene Seymour'))
