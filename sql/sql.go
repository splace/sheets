package table

import "regexp"

const label ="[a-zA-Z][a-zA-Z0-9_]*"

var sqlre = regexp.MustCompile(`^\s*(SELECT|INSERT INTO|DELETE|UPDATE)\s+(\*|`+label+`]*)\s+FROM\s+(`+label+`)\s*WHERE\s+(.+?)\s+(?:LIMIT\s+([1-9][0-9]*))?`)


//strings.Fields
//slices.BinarySearch

var resWords = []string{"ADD","ADD CONSTRAINT","ALL","ALTER","ALTER COLUMN","ALTER TABLE","AND","ANY","AS","ASC","BACKUP DATABASE","BETWEEN","CASE","CHECK","COLUMN","CONSTRAINT","CREATE","CREATE DATABASE","CREATE INDEX","CREATE OR REPLACE VIEW","CREATE TABLE","CREATE PROCEDURE","CREATE UNIQUE INDEX","CREATE VIEW","DATABASE","DEFAULT","DELETE","DESC","DISTINCT","DROP","DROP COLUMN","DROP CONSTRAINT","DROP DATABASE","DROP DEFAULT","DROP INDEX","DROP TABLE","DROP VIEW","EXEC","EXISTS","FOREIGN KEY","FROM","FULL OUTER JOIN","GROUP BY","HAVING","IN","INDEX","INNER JOIN","INSERT INTO","INSERT INTO SELECT","IS NULL","IS NOT NULL","JOIN","LEFT JOIN","LIKE","LIMIT","NOT","NOT NULL","OR","ORDER BY","OUTER JOIN","PRIMARY KEY","PROCEDURE","RIGHT JOIN","ROWNUM","SELECT","SELECT DISTINCT","SELECT INTO","SELECT TOP","SET","TABLE","TOP","TRUNCATE TABLE","UNION","UNION ALL","UNIQUE","UPDATE","VALUES","VIEW","WHERE"}

//ADD 	Adds a column in an existing table
//ADD CONSTRAINT 	Adds a constraint after a table is already created
//ALL 	Returns true if all of the subquery values meet the condition
//ALTER 	Adds, deletes, or modifies columns in a table, or changes the data type of a column in a table
//ALTER COLUMN 	Changes the data type of a column in a table
//ALTER TABLE 	Adds, deletes, or modifies columns in a table
//AND 	Only includes rows where both conditions is true
//ANY 	Returns true if any of the subquery values meet the condition
//AS 	Renames a column or table with an alias
//ASC 	Sorts the result set in ascending order
//BACKUP DATABASE 	Creates a back up of an existing database
//BETWEEN 	Selects values within a given range
//CASE 	Creates different outputs based on conditions
//CHECK 	A constraint that limits the value that can be placed in a column
//COLUMN 	Changes the data type of a column or deletes a column in a table
//CONSTRAINT 	Adds or deletes a constraint
//CREATE 	Creates a database, index, view, table, or procedure
//CREATE DATABASE 	Creates a new SQL database
//CREATE INDEX 	Creates an index on a table (allows duplicate values)
//CREATE OR REPLACE VIEW 	Updates a view
//CREATE TABLE 	Creates a new table in the database
//CREATE PROCEDURE 	Creates a stored procedure
//CREATE UNIQUE INDEX 	Creates a unique index on a table (no duplicate values)
//CREATE VIEW 	Creates a view based on the result set of a SELECT statement
//DATABASE 	Creates or deletes an SQL database
//DEFAULT 	A constraint that provides a default value for a column
//DELETE 	Deletes rows from a table
//DESC 	Sorts the result set in descending order
//DISTINCT 	Selects only distinct (different) values
//DROP 	Deletes a column, constraint, database, index, table, or view
//DROP COLUMN 	Deletes a column in a table
//DROP CONSTRAINT 	Deletes a UNIQUE, PRIMARY KEY, FOREIGN KEY, or CHECK constraint
//DROP DATABASE 	Deletes an existing SQL database
//DROP DEFAULT 	Deletes a DEFAULT constraint
//DROP INDEX 	Deletes an index in a table
//DROP TABLE 	Deletes an existing table in the database
//DROP VIEW 	Deletes a view
//EXEC 	Executes a stored procedure
//EXISTS 	Tests for the existence of any record in a subquery
//FOREIGN KEY 	A constraint that is a key used to link two tables together
//FROM 	Specifies which table to select or delete data from
//FULL OUTER JOIN 	Returns all rows when there is a match in either left table or right table
//GROUP BY 	Groups the result set (used with aggregate functions: COUNT, MAX, MIN, SUM, AVG)
//HAVING 	Used instead of WHERE with aggregate functions
//IN 	Allows you to specify multiple values in a WHERE clause
//INDEX 	Creates or deletes an index in a table
//INNER JOIN 	Returns rows that have matching values in both tables
//INSERT INTO 	Inserts new rows in a table
//INSERT INTO SELECT 	Copies data from one table into another table
//IS NULL 	Tests for empty values
//IS NOT NULL 	Tests for non-empty values
//JOIN 	Joins tables
//LEFT JOIN 	Returns all rows from the left table, and the matching rows from the right table
//LIKE 	Searches for a specified pattern in a column
//LIMIT 	Specifies the number of records to return in the result set
//NOT 	Only includes rows where a condition is not true
//NOT NULL 	A constraint that enforces a column to not accept NULL values
//OR 	Includes rows where either condition is true
//ORDER BY 	Sorts the result set in ascending or descending order
//OUTER JOIN 	Returns all rows when there is a match in either left table or right table
//PRIMARY KEY 	A constraint that uniquely identifies each record in a database table
//PROCEDURE 	A stored procedure
//RIGHT JOIN 	Returns all rows from the right table, and the matching rows from the left table
//ROWNUM 	Specifies the number of records to return in the result set
//SELECT 	Selects data from a database
//SELECT DISTINCT 	Selects only distinct (different) values
//SELECT INTO 	Copies data from one table into a new table
//SELECT TOP 	Specifies the number of records to return in the result set
//SET 	Specifies which columns and values that should be updated in a table
//TABLE 	Creates a table, or adds, deletes, or modifies columns in a table, or deletes a table or data inside a table
//TOP 	Specifies the number of records to return in the result set
//TRUNCATE TABLE 	Deletes the data inside a table, but not the table itself
//UNION 	Combines the result set of two or more SELECT statements (only distinct values)
//UNION ALL 	Combines the result set of two or more SELECT statements (allows duplicate values)
//UNIQUE 	A constraint that ensures that all values in a column are unique
//UPDATE 	Updates existing rows in a table
//VALUES 	Specifies the values of an INSERT INTO statement
//VIEW 	Creates, updates, or deletes a view
//WHERE	Filters a result set to include only records that fulfill a specified condition


//1. Data Definition Language (DDL)
//Data Definition Language (DDL) commands are used for creating a table, deleting a table, altering a table, and truncating the table. All the commands of DDL are auto-committed i.e. it permanently save all the changes in the database.

//In this category, we have six commands:

//Command	Description	Example
//CREATE	Used to create new database objects such as tables, databases, or indexes.	CREATE TABLE Students ( ID INT, Name VARCHAR(50), Age INT );
//ALTER	Modifies the structure of an existing table, such as adding, deleting, or modifying columns.	ALTER TABLE Students ADD Email VARCHAR(100);
//DROP	Deletes an existing database object, such as a table or database.	DROP TABLE Students;
//TRUNCATE	Removes all records from a table without logging individual row deletions, but keeps the table structure.	TRUNCATE TABLE Students;
//RENAME	Renames a database object, such as a table or column.	RENAME TABLE Students TO Learners;
//COMMENT	Adds comments or annotations to the database objects.	COMMENT ON TABLE Students IS 'Table for storing student informati



//Data Manipulation Language (DML)
//These SQL commands modify the database. The commands of DML are not auto-committed i.e. it can't permanently save all the changes in the database. They can be rolled back. In this category, we have the following commands.

//Command	Description	Syntax Example
//INSERT	Adds new rows of data into a table.	INSERT INTO table_name (column1,column2) VALUES (value1, value2);
//UPDATE	Modifies existing data in a table.	UPDATE table_name SET column1 = value1 WHERE condition;
//DELETE	Removes existing data from a table based on a condition.	DELETE FROM table_name WHERE condition;
//MERGE	Combines INSERT and UPDATE operations based on a condition (supported in some SQL implementations).	MERGE INTO table_name USING source_table ON condition WHEN MATCHED THEN UPDATE WHEN NOT MATCHED THEN INSERT;




//3. Data Query Language (DQL)
//This SQL command is used to fetch/retrieve data from database tables.

//Command	Description	Example
//SELECT	Used to query and retrieve data from one or more tables.	SELECT * FROM Students;
//SELECT DISTINCT	Retrieves unique values from a column, avoiding duplicates.	SELECT DISTINCT Department FROM Employees;
//WHERE	Filters rows based on specified conditions.	SELECT * FROM Students WHERE Age > 18;
//ORDER BY	Sorts the result set based on one or more columns in ascending or descending order.	SELECT * FROM Employees ORDER BY Salary DESC
//GROUP BY	Group rows share a property and perform aggregate functions on each group.	SELECT Department, COUNT(*) FROM Employees GROUP BY Department;
//HAVING	Filters grouped data based on conditions, used with GROUP BY.	SELECT Department, AVG(Salary) FROM Employees GROUP BY Department HAVING AVG(Salary) > 50000;
//LIMIT	Restricts the number of rows returned in the result.	SELECT * FROM Products LIMIT 10;
//JOIN	Combines rows from two or more tables based on a related column.	SELECT Orders.OrderID, Customers.Name FROM Orders JOIN Customers ON Orders.CustomerID = Customers.ID




//Transaction Control Language (TCL)
//These SQL commands are used to handle changes affecting the data in the database. We use these commands within the transaction or to make a stable point during changes in the database, at which we can roll back the database state if required.

//Command	Description	Example
//COMMIT	Saves all the transactions permanently to the database.	COMMIT;
//ROLLBACK	Undoes transactions that are not yet saved to the database.	ROLLBACK;
//SAVEPOINT	Sets a point within a transaction to which you can later roll back.	SAVEPOINT savepoint_name;
//RELEASE SAVEPOINT	Deletes a savepoint, making it unavailable for future rollbacks.	RELEASE SAVEPOINT savepoint_name;
//SET TRANSACTION	Defines a transaction's characteristics, such as isolation level or read/write.	SET TRANSACTION ISOLATION LEVEL SERIALI




//Data Control Language (DCL)
//Data Control Language (DCL) commands are used to implement security on database objects like tables, views, stored procedures, etc. It consists of commands that deal with the user permissions and controls of the database system. In this category, we have two main commands:

//1. GRANT
//The GRANT command in SQL is used to give specific permissions to users, such as access to read, insert, or update data in a table. It helps manage and control who can do what in the database.

//.
//2. REVOKE
