from critics import critics
from pearson import sim_pearson

def get_recommendations(prefs, person, similarity=sim_pearson):
    totals = {}
    sim_sums = {}

    for other in prefs:
        # Don't include self
        if other == person: continue

        sim = similarity(prefs, person, other)

        # Ignore scores of 0 or lower
        if sim <= 0: continue

        for item in prefs[other]:
            if item not in prefs[person] or prefs[person][item] == 0:
                # similarity * score
                totals.setdefault(item, 0)
                totals[item] += prefs[other][item] * sim

                # Sum of similarities
                sim_sums.setdefault(item, 0)
                sim_sums[item] += sim

    rankings = [(total / sim_sums[item], item) for item, total in totals.items()]

    rankings.sort()
    rankings.reverse()
    return rankings

print get_recommendations(critics, 'Toby')
