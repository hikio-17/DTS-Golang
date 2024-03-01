<?php
function run($x, $y) {
$data = [
	[
		"nama" => "Paket Harmoni",
		"harga" => 149000
	],
	[
		"nama" => "Paket Ritme",
		"harga" => 219000
	],
    [
		"nama" => "Paket Melodi",
		"harga" => 299000
	],
    [
		"nama" => "Paket Akustik",
		"harga" => 99000
	],
    [
		"nama" => "Paket Beat",
		"harga" => 379000
	],
    [
		"nama" => "Paket Harmony Plus",
		"harga" => 449000
	],
    [
		"nama" => "Paket Duet",
		"harga" => 189000
	],
    [
		"nama" => "Paket Symphony",
		"harga" => 599000
	],
    [
		"nama" => "Paket Tempo",
		"harga" => 259000
	],
    [
		"nama" => "Paket Rhythm Plus",
		"harga" => 499000
	],
    
    
];

$totalPrice = $data[$x]["harga"] * $y;

echo $totalPrice;
}

// run(0, 2);

function run1 ($k, $c) {
	$countChar = substr_count($k, $c);
    echo $countChar;
}

// run1("programmer", "r");

function convertKupon($n) {
	return 100 * (2 ** ($n -1));
}

function run2($n, $u) {
	$kupon;
	$returnPrice;
	if ($n > $u) {
    	$returnPrice = 0;
    } else if ($n > 2000000 && $n < 5000000) {
    	$kupon = 7;
    	$returnPrice = $u - ($n - (5/100 * $n));
    } else if ($n > 5000000 && $n < 10000000) {
    	$kupon = 10;
    	$returnPrice = $u - ($n - (10/100 * $n));
    } else if ($n >= 10000000) {
    	$kupon = 12;
    	$returnPrice = $u - ($n - (20/100 * $n));
    } else {
    	$kupon = 5;
    	$returnPrice = $u - $n;
    }
    
    $returnPriceWithKupon = $returnPrice + convertKupon($kupon);
    
    echo $returnPriceWithKupon;
    return $returnPriceWithKupon;
}

run2(7550000, 9000000);
?>