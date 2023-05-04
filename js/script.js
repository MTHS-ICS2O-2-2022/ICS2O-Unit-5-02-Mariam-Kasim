// Copyright (c) 2022 Mariam Kasim All rights reserved
//
// Created by: Mariam Kasim
// Created on: Apr 26
// This file contains the JS functions for index.html

"use strict"

// Function to generate random potive and negative numbers
function myButtonClicked() {
  // Generate random number
  let randomNumber = 0.00

  // input
  const positiveButtonChecked = document.getElementById('positive').checked

  // process
  if (positiveButtonChecked == true) {
    // positive
    randomNumber = Math.floor(Math.random() * 6)
  } else {
    // negative
    randomNumber = Math.floor(Math.random() * -6)
  }

  // output
  document.getElementById('answer').innerHTML = "The random number is: " + randomNumber
}
