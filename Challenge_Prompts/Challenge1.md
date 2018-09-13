# Challenge 1: Identifying COLREGs Interactions 

Out of the millions of data points in the AIS dataset can you find the most efficient algorithm to identify potential COLREGs interactions?

COLREGs interactions are defined by situations where there are constant bearing and decreasing range between two ships. In other words, two ships are headed for collision unless something changes. COLREGs also assume that two ships are visually in sight of one another. You can assume they are a maximum of 4 nautical miles away. This challenge is about developing a screening tool for COLREGs interactions within AIS dataset, which is akin to finding the proverbial needle in the haystack. 
As a starting point, consider looking for places where ships come close to each other. Ships frequently come close to each other in high traffic areas, like ports, but often the ships are not moving, or the interactions are governed by prescriptive criteria, like traffic lanes.

The more interesting interactions for understanding COLREGs compliance are in areas where the traffic patterns have not explicitly defined. 
You may use any of the AIS data to address this challenge, but we would suggest looking at a high traffic area with rules, like San Diego or the Puget Sound (UTM 10 and 11) and an area that has less rules and more open ocean, like The Caribbean (UTM 17 and 18, less than 25 degrees latitude), as a starting point.  
From there you can refine your algorithms in any way you see fit. You will find that many of the instances where ships are in close proximity for an extended duration are activities such as underway supply or towing into port and the results of this challenge will be improved by understanding the variables discussed in Challenge 2.

The judging criteria for the challenge will include: 
- Computational efficiency: 40 points. 
- Accuracy: 40 points. Do the interactions you found actually represent COLREGs encounters or just places where ships were in close proximity?
- Data presentation and visualization: 20 points.
