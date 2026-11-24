# Repeating and Missing Number

> Striver SDE Sheet 011 · Day 2 Arrays — XOR everything down to repeated^missing, then split by the lowest set bit. Go, stdlib only, tested.

## Problem

An array of size n was supposed to contain 1..n exactly once, but one value appears twice and another is absent. Find both — the repeated and the missing — in O(n) time and O(1) space, without modifying the array.

## Approach

XOR all array values together with all numbers 1..n: paired values cancel, leaving `x = repeated ^ missing`. Any set bit of x (take the lowest, `x & -x`) is a bit where the two differ — so partition every array value *and* every number 1..n into two XOR buckets by that bit. Each bucket now cancels down to one of the two answers. A final scan checks which candidate actually occurs in the array, labeling repeated vs missing. Unlike the sum/sum-of-squares method, XOR never overflows.

**Complexity:** O(n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
