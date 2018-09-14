# TRACK 2: Data Science and the Seven Seas: Collision Avoidance

<p align="center">
  <img src="https://static1.squarespace.com/static/596d24cd4402430bb863ffad/t/5b41e62603ce641f98f2e3cd/1536741696061/?format=1500w" width="350" title="hover text">
</p>


International regulations for preventing collisions at sea (COLREGS) codify long-established international norms for maritime navigation. Mariners are held accountable to the principles in COLREGS for safely handling their vessels. However, many situations that are frequently encountered by mariners, such as three vessels converging, are not covered by COLREGS. Similar to automotive traffic laws, the safest maneuver, in a given situation, depends on many factors.  For example, turning right may be legal in two different situations.  At an intersection on a country road, this maneuver would be safe and failing to turn would result in an angry train of drivers behind you.  If you tried that same right turn in Times Square in NYC you would find yourself trapped in a swarm of pedestrians, unable to safely make the turn.  The context-dependent nature of COLREGS makes it particularly difficult for autonomous systems, which lack human judgement. This challenge will use data from ships underway on the high seas to develop algorithms to assist the Navy with preventing collisions for human-operated and autonomous vessels. 

## The Data:

For this challenge, your primary data is collected from the Automatic Identification System (AIS). Ships carry AIS to provide unique identification, position, course, and speed data to other ships and shore stations. AIS is intended to assist a vessel's crew with safe navigation and maritime authorities with monitoring vessel movements. The International Maritime Organization requires AIS to be fitted aboard international voyaging ships with 300 or more gross tonnage (GT), and all passenger ships regardless of size.

AIS data is enormous and rich (the 2016 United States dataset alone has 9.4 billion records), but it was designed as a safety feature for local collision avoidance communication between ships at sea, not to be analyzed holistically. As such AIS data provides huge opportunity for understanding the contextual nature of COLREGS encounters for ships at sea, but it also presents significant challenges including:
- Noise: AIS systems transmit radio signals in the Maritime VHF band, which is inherently noisy. 
- Transmission Gaps: There are often jumps in the data due to spotty AIS signals, low density of receivers, or vessel operators turning off the signal to avoid detection.
- Congestion: In crowded areas, bandwidths become congested and signals interfere with one another. 
- Errors: Sometimes ships appear in impossible places like on land or airports or vessel operators mis-enter AIS codes. This might be a mistake or malfunction or can be indicative of malicious activity (https://www.wired.co.uk/article/black-sea-ship-hacking-russia).

Determining “normal” behavior is incredibly difficult and something you can help the US Navy tackle through the following 2 challenges. It is expected that most teams will work on both challenges and judging will be a cumulative total.

<p align="center">
  <img src="https://static.wixstatic.com/media/3d35e8_2d9eb95a4abe4869afafbf51d29038dc~mv2.png/v1/fill/w_288,h_60,al_c,usm_0.66_1.00_0.01/3d35e8_2d9eb95a4abe4869afafbf51d29038dc~mv2.png" width="150" title="hover text">
</p>
