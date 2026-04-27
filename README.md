# How does Go compare to Python and R with Linear Regression

## Overview 

**Can Go produce regression results comparable to Python and R?** 

This project tests Go statistics package (montanaflynn/stats) by comparing linear regression results on the Anscombe Quartet with Python and R.


### Results

The Go statistics package outputs matched those of R and Python. The execution time was significantly faster. All three implementations produced nearly identical regression coefficients, confirming the correctness of the Go package. 

#### Validated Statistics

- Intercept
- Slope
- R-squared
- Standard Error
- t-statistic
- p-value
- F-statistic
- Adjusted R-squared

#### Set 1 Results

| Statistic | Go | Python | R |
|-----------|-----|--------|---|
| Intercept | 3.0001 | 3.0001 | 3.0001 |
| Slope | 0.5001 | 0.5001 | 0.5001 |
| R-squared | 0.6665 | 0.667 | 0.6665 |
| Std Error | 0.1179 | 0.118 | 0.1179 |
| t-statistic | 4.2415 | 4.241 | 4.241 |
| p-value | 0.0022 | 0.002 | 0.00217 |
| F-statistic | 17.9899 | 17.99 | 17.99 |
| Adj R-squared | 0.6295 | 0.629 | 0.6295 |

The additional test results are in the .txt files: 
- `go_results.txt`
- `python_results.txt`
- `r_results.txt`

### Performance & Memory Comparison
Batch execution benchmarks revealed the following performance differences:

| Language | Execution Time | Peak Memory Usage |
| :--- | :--- | :--- |
| **Go** | 0.000067 seconds | Low |
| **R** | 0.00835 seconds | High |
| **Python** | 0.0327 seconds | High |

Go executed significantly faster than both Python and R. Memory measurements are approximate because Go, Python, and R report memory differently. Go reported the lowest memory use in this small test, but the results should be interpreted as general indicators rather than exact one-to-one comparisons.

### Recommendation 

Based on these results, Go is capable of producing correct statistical outputs and offers strong performance advantages. However, data scientists may have concerns about the limited statistical libraries compared to Python and R. Go provides accurate results with currently available packages, but programmers may encounter gaps in specialized statistical methods, visualization tools, and community support. 

Suggested Approach:

Use Go for backend, deployment, and performance-sensitive statistical utilities, while maintaining Python and R for exploratory data analysis, visualization, and advanced statistical modeling.


### Getting Started 

- `figures` - Figures output from the program
    - `fig_anscombe_Python.pdf` - Scatter plot visualizations generated via Python.
    - `fig_anscombe_R.pdf` - Scatter plot visualizations generated via R.
- `go.mod` 
- `go.sum` 
- `main.go` - Main Go program 
- `main_test.go` - Go testing used to check the program
- `miller-mtpa-chapter-1-program.py` - Python Program
- `miller-mtpa-chapter-1-program.R` - R Program
- `README.md` 
- `results`- Results output from the program 
    - `go_results.txt` - Output for Go with runtime
    - `python_results.txt` - Output for Python with runtime
    - `r_results.txt` - Output for R with runtime

### Run with Go

```bash
go run main.go
```

### Run with R

```bash
Rscript miller-mtpa-chapter-1-program.R
```

### Run with Python 
```bash
pip install pandas numpy statsmodels matplotlib
python miller-mtpa-chapter-1-program.py
```

### Build an Executable

#### Windows

```powershell
go build -o main.exe main.go
.\main.exe
```

#### MacOS / Linux

```bash
go build -o main main.go
./main
```

#### Testing 

```bash
go test
```

#### Benchmark 

```bash
go test -bench=. -benchmem 
```

### Creating the files with result 
(this does not put them in folders ) 

```
go run main.go > go_results.txt
python miller-mtpa-chapter-1-program.py > python_results.txt
Rscript miller-mtpa-chapter-1-program.R > r_results.txt
```

### AI Disclosure 

This project used GenAI as a supplemental tool for README refinement, code review, debugging support, and clarification of Go tooling and syntax. 

**Notes on AI**

It was helpful for building understanding but often suggested more complicated versions. 