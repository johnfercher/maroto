# PDF417 barcodes in Golang

[![Build Status](https://travis-ci.org/ruudk/golang-pdf417.svg?branch=master)](https://travis-ci.org/ruudk/golang-pdf417)

This is a port of https://github.com/ihabunek/pdf417-php

This library encodes data to a PixelGrid that can be used to display the barcode.
You can use the PixelGrid to draw the barcode on anything. Check [pdf417_test.go](pdf417_test.go) for an example.

I only needed a way to draw PDF417 barcodes onto a [FPDF](https://github.com/jung-kurt/gofpdf) PDF file. 
See example: [examples/fpdf.go](examples/fpdf.go)

If you want to export the barcode to an image, you have to create something that puts the PixelGrid onto the image. If you find a way, please submit a PR :)

