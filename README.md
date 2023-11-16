# golang-and-redis
1. The "main.go" has a producer that generates objects of a structure person. <br />
2. The person has 2 fields: Name and Age.
3. Each object created is assigned a random name from the slice name that consist of multiple random names
4. Each object is given a random age using the GenerateRandomAge function, which ensures the age falls within the range of 30 to 50.
5. The objects are pushed into a Redis queue and stored in an in-memory list.
6. Established a connection pool to minimize the time required for insertion.
7. Implemented a switch function that allows users to select the desired quantity of objects to generate and store in Redis.
8. A menu is desplayed with choices, For example, entering '1' produces 100,000 objects, '2' produces 200,000 objects, '3' produces 300,000 objects, and so forth.
<br/>

**Benchmarks:** <br /><br />
1. To insert 100,000 objects in the database, the createAndStoreObjectsInQueue function takes 1,443,780 (1.4 million) microseconds on average, equivalent to 1.44378 seconds.
2. To insert 200,000 objects in the database, the createAndStoreObjectsInQueue function takes 2,917,780 (2.9 million) microseconds on average, equivalent to 2.91778 seconds.
3. To insert 300,000 objects in the database, the createAndStoreObjectsInQueue function takes 4,549,820 (4.5 million) microseconds on average, equivalent to 4.54982 seconds.
4. To insert 400,000 objects in the database, the createAndStoreObjectsInQueue function takes 5,916,710 (5.9 million) microseconds on average, equivalent to 5.91671 seconds.
5. To insert 500,000 objects in the database, the createAndStoreObjectsInQueue function takes 7,331,480 (7.3 million) microseconds on average, equivalent to 7.33148 seconds.

