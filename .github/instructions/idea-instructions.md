# Overview
I want to create a simple web app as below

- language: go1.23.5
- framework: echo

Feature: create a printable calendar for users to print and track their goal progress.

# Screens:

The overall UI color should be back and white.

Please make a simple UI but having some sort of good-looking.

It should display well on both smartphone and PC and tablet.

## URI: GET goals/add

This is an input screen.

Input items:
  - Goal name: as long as 100 characters
  - Goal description: as long as 255 characters
  - Start
  - End: must be a future day which as no more than 25 days from the `start`

It has submit and cancel button. On submitting, save data into browser storage

Internally assign an id for the goals