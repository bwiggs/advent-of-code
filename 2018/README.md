# Day 7 - Topological Sort

Day 7 messed me up. My brain really struggled to solve this and it lead to a lot of mental frustration that I couldn't come up with a clean solution. Not that I couldn't code up a solution, but jsut that the way I was going to solve it felt really dirty.

So I cheated and looked at how some others were solving it. That's when I learned about Topological Sorts. Topo sorts was new to me.

After lookup up Topo sorts and reading other's code I found out that I had misread the requirements. I thought that all child steps had to be completed first. Well friends, if you reread the instructions, you'll see that in reality you just need to process the next alphabetical step in the list, regardless of it's location in the graph. THAT MAKES THINGS WAY EASIER.

Thoughts on GoLang:

1. **GoLang is verbose**. Setting up maps and filling them requires a lot of nil checking. Nothing is free.
2. **List processesing isnt as easy** as it is in functional languages. Being able to slice and dice arrays is really convinent, and probably one of the reasons why a good amount of AoC participants are using Python to solve these problems.

