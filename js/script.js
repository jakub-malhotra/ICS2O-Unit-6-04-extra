// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Jakub Malhotra
// Created on: May 2023
// This file contains the JS functions for index.html

"use strict"

window.onload = function () {
  // this calculates fahrenheit to celsius
  const params = new URLSearchParams(document.location.search)
  const fahrenheit = params.get("f")

  document.getElementById("fahrenheit").innerHTML =
    "The temperature in fahrenheit is: " + fahrenheit + "℉."

  const celsius = ((fahrenheit - 32) * 5) / 9
  document.getElementById("celsius").innerHTML =
    "The temperature in celsius is: " + celsius.toFixed(2) + "℃."
}