from critics import critics
from pearson import sim_pearson

def top_matches(prefs, person, n=5, similarity=sim_pearson):
    scores = [(similarity(prefs, person, other), other) for other in prefs if other !=person]

    scores.sort()
    scores.reverse()
    return scores[0:n]

print top_matches(critics, 'Toby', n=3)
