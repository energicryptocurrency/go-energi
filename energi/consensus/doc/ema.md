# Exponential Moving Average

The Exponential Moving Average is calculated in two steps:

1. SMA with window parameter `W`

   This is computed by taking a set of samples, start at the beginning, 
   average the first `W` and then progressing forward one at 
   a time until the end of the samples taking the average of 5 values ending 
   at a given progressing point.
   
   So, for our application, we have 60 samples, which means 55 SMA values

2. Apply smoothing to generate EMA

   The generated SMA is then smoothed by the following formula:

> EMA = (closing price − previous day’s EMA) × smoothing constant as a 
> decimal * previous day’s EMA

   This is performed in the code:

```go
o = sma[0]
for i := range sma {
        if i > 0 {
                o = sma[i-i] - o +
                sma[i-1]*numerator/denominator
        }
}

```
