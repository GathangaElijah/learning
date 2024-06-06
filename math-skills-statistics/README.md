# math-skills

### Project Description
This repository contains a program that calculates; <b>average</b>, <b>median</b>, <b>variance</b> and<b> standard deviation</b>. 

# Running the program

### Requirements

#### **Note:** ***(This project was created on ubuntu linux)***

To ensure the program run correctly provide a file `data.txt` which contains the data. This data should appear in this format;

```
8
4
6
3
3
5
2
1
9
7
11
8
33
54
```

The program runs on <b> go version 1.22.1</b> or higher. To run the program follow the following steps;
1. Ensure Go is installed [how to install go](https://go.dev/doc/install)
2. Clone the repository  

```
git clone https://learn.zone01kisumu.ke/git/egathang/math-skills.git
```

3. Navigate to the program directory 

```
cd math-skills
```

4. Copy the data you have in your file to `data.txt` file in the program. ***(optional)***

    **Note:** `data.txt` file contains default data.
    If you want to use your own data, make sure you open the program in a code editor like **vscode** and paste your data in the program's data.txt file.

5. Run the program 

```
go run . "data.txt"
```

When you run the program you expect an output that looks like this:

```
Average: 11
Median: 7
Variance: 199
Standard Deviation: 14
```

