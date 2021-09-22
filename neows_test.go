package nasa

import (
	"encoding/json"
	"reflect"
	"testing"
)

var mockData = `{
	"links": {
	"next": "http://www.neowsapp.com/rest/v1/feed?start_date=2015-09-08&end_date=2015-09-09&detailed=false&api_key=DEMO_KEY",
	"prev": "http://www.neowsapp.com/rest/v1/feed?start_date=2015-09-06&end_date=2015-09-07&detailed=false&api_key=DEMO_KEY",
	"self": "http://www.neowsapp.com/rest/v1/feed?start_date=2015-09-07&end_date=2015-09-08&detailed=false&api_key=DEMO_KEY"
	},
	"element_count": 24,
	"near_earth_objects": {
	"2015-09-08": [
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/2465633?api_key=DEMO_KEY"
	},
	"id": "2465633",
	"neo_reference_id": "2465633",
	"name": "465633 (2009 JR5)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=2465633",
	"absolute_magnitude_h": 20.36,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.2251930467,
	"estimated_diameter_max": 0.5035469604
	},
	"meters": {
	"estimated_diameter_min": 225.1930466786,
	"estimated_diameter_max": 503.5469604336
	},
	"miles": {
	"estimated_diameter_min": 0.1399284286,
	"estimated_diameter_max": 0.3128894784
	},
	"feet": {
	"estimated_diameter_min": 738.8223552649,
	"estimated_diameter_max": 1652.0570096689
	}
	},
	"is_potentially_hazardous_asteroid": true,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-08",
	"close_approach_date_full": "2015-Sep-08 20:28",
	"epoch_date_close_approach": 1441744080000,
	"relative_velocity": {
	"kilometers_per_second": "18.1279547773",
	"kilometers_per_hour": "65260.6371983344",
	"miles_per_hour": "40550.4220413761"
	},
	"miss_distance": {
	"astronomical": "0.3027478814",
	"lunar": "117.7689258646",
	"kilometers": "45290438.204452618",
	"miles": "28142173.3303294084"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3426410?api_key=DEMO_KEY"
	},
	"id": "3426410",
	"neo_reference_id": "3426410",
	"name": "(2008 QV11)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3426410",
	"absolute_magnitude_h": 21.34,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.1434019235,
	"estimated_diameter_max": 0.320656449
	},
	"meters": {
	"estimated_diameter_min": 143.4019234645,
	"estimated_diameter_max": 320.6564489709
	},
	"miles": {
	"estimated_diameter_min": 0.0891057966,
	"estimated_diameter_max": 0.1992466184
	},
	"feet": {
	"estimated_diameter_min": 470.4787665793,
	"estimated_diameter_max": 1052.0225040417
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-08",
	"close_approach_date_full": "2015-Sep-08 14:31",
	"epoch_date_close_approach": 1441722660000,
	"relative_velocity": {
	"kilometers_per_second": "19.7498128142",
	"kilometers_per_hour": "71099.3261312856",
	"miles_per_hour": "44178.3562841869"
	},
	"miss_distance": {
	"astronomical": "0.2591250701",
	"lunar": "100.7996522689",
	"kilometers": "38764558.550560687",
	"miles": "24087179.7459520006"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3553060?api_key=DEMO_KEY"
	},
	"id": "3553060",
	"neo_reference_id": "3553060",
	"name": "(2010 XT10)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3553060",
	"absolute_magnitude_h": 26.5,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.0133215567,
	"estimated_diameter_max": 0.0297879063
	},
	"meters": {
	"estimated_diameter_min": 13.3215566698,
	"estimated_diameter_max": 29.7879062798
	},
	"miles": {
	"estimated_diameter_min": 0.008277629,
	"estimated_diameter_max": 0.0185093411
	},
	"feet": {
	"estimated_diameter_min": 43.7058959846,
	"estimated_diameter_max": 97.7293544391
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-08",
	"close_approach_date_full": "2015-Sep-08 12:07",
	"epoch_date_close_approach": 1441714020000,
	"relative_velocity": {
	"kilometers_per_second": "19.1530348886",
	"kilometers_per_hour": "68950.9255988812",
	"miles_per_hour": "42843.4237422604"
	},
	"miss_distance": {
	"astronomical": "0.4917435147",
	"lunar": "191.2882272183",
	"kilometers": "73563782.385433689",
	"miles": "45710414.7542113482"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3726710?api_key=DEMO_KEY"
	},
	"id": "3726710",
	"neo_reference_id": "3726710",
	"name": "(2015 RC)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3726710",
	"absolute_magnitude_h": 24.3,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.0366906138,
	"estimated_diameter_max": 0.0820427065
	},
	"meters": {
	"estimated_diameter_min": 36.6906137531,
	"estimated_diameter_max": 82.0427064882
	},
	"miles": {
	"estimated_diameter_min": 0.0227984834,
	"estimated_diameter_max": 0.0509789586
	},
	"feet": {
	"estimated_diameter_min": 120.3760332259,
	"estimated_diameter_max": 269.1689931548
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-08",
	"close_approach_date_full": "2015-Sep-08 09:45",
	"epoch_date_close_approach": 1441705500000,
	"relative_velocity": {
	"kilometers_per_second": "19.486643553",
	"kilometers_per_hour": "70151.9167909206",
	"miles_per_hour": "43589.6729637806"
	},
	"miss_distance": {
	"astronomical": "0.0269252677",
	"lunar": "10.4739291353",
	"kilometers": "4027962.697099799",
	"miles": "2502859.9608192662"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3727181?api_key=DEMO_KEY"
	},
	"id": "3727181",
	"neo_reference_id": "3727181",
	"name": "(2015 RO36)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3727181",
	"absolute_magnitude_h": 22.9,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.0699125232,
	"estimated_diameter_max": 0.1563291544
	},
	"meters": {
	"estimated_diameter_min": 69.9125232246,
	"estimated_diameter_max": 156.3291544087
	},
	"miles": {
	"estimated_diameter_min": 0.0434416145,
	"estimated_diameter_max": 0.097138403
	},
	"feet": {
	"estimated_diameter_min": 229.3718026961,
	"estimated_diameter_max": 512.8909429502
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-08",
	"close_approach_date_full": "2015-Sep-08 14:36",
	"epoch_date_close_approach": 1441722960000,
	"relative_velocity": {
	"kilometers_per_second": "15.8053596703",
	"kilometers_per_hour": "56899.294813224",
	"miles_per_hour": "35355.0090465835"
	},
	"miss_distance": {
	"astronomical": "0.0540392535",
	"lunar": "21.0212696115",
	"kilometers": "8084157.219990045",
	"miles": "5023262.364730821"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3727639?api_key=DEMO_KEY"
	},
	"id": "3727639",
	"neo_reference_id": "3727639",
	"name": "(2015 RN83)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3727639",
	"absolute_magnitude_h": 21.7,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.1214940408,
	"estimated_diameter_max": 0.2716689341
	},
	"meters": {
	"estimated_diameter_min": 121.4940407996,
	"estimated_diameter_max": 271.6689340891
	},
	"miles": {
	"estimated_diameter_min": 0.0754928736,
	"estimated_diameter_max": 0.1688071972
	},
	"feet": {
	"estimated_diameter_min": 398.6025088171,
	"estimated_diameter_max": 891.3023057169
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-08",
	"close_approach_date_full": "2015-Sep-08 15:42",
	"epoch_date_close_approach": 1441726920000,
	"relative_velocity": {
	"kilometers_per_second": "12.0811420305",
	"kilometers_per_hour": "43492.1113096542",
	"miles_per_hour": "27024.3066079349"
	},
	"miss_distance": {
	"astronomical": "0.1684193589",
	"lunar": "65.5151306121",
	"kilometers": "25195177.358205543",
	"miles": "15655557.2525527734"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3730577?api_key=DEMO_KEY"
	},
	"id": "3730577",
	"neo_reference_id": "3730577",
	"name": "(2015 TX237)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3730577",
	"absolute_magnitude_h": 23.3,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.058150704,
	"estimated_diameter_max": 0.130028927
	},
	"meters": {
	"estimated_diameter_min": 58.1507039646,
	"estimated_diameter_max": 130.0289270043
	},
	"miles": {
	"estimated_diameter_min": 0.0361331611,
	"estimated_diameter_max": 0.0807962044
	},
	"feet": {
	"estimated_diameter_min": 190.7831555951,
	"estimated_diameter_max": 426.6041048727
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-08",
	"close_approach_date_full": "2015-Sep-08 14:19",
	"epoch_date_close_approach": 1441721940000,
	"relative_velocity": {
	"kilometers_per_second": "6.573400491",
	"kilometers_per_hour": "23664.2417675863",
	"miles_per_hour": "14704.0395583094"
	},
	"miss_distance": {
	"astronomical": "0.0795238758",
	"lunar": "30.9347876862",
	"kilometers": "11896602.433824546",
	"miles": "7392205.9712328948"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3731587?api_key=DEMO_KEY"
	},
	"id": "3731587",
	"neo_reference_id": "3731587",
	"name": "(2015 UG)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3731587",
	"absolute_magnitude_h": 22.81,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.0728710415,
	"estimated_diameter_max": 0.1629446024
	},
	"meters": {
	"estimated_diameter_min": 72.8710414898,
	"estimated_diameter_max": 162.9446023625
	},
	"miles": {
	"estimated_diameter_min": 0.0452799519,
	"estimated_diameter_max": 0.1012490505
	},
	"feet": {
	"estimated_diameter_min": 239.0782277615,
	"estimated_diameter_max": 534.595169215
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-08",
	"close_approach_date_full": "2015-Sep-08 18:50",
	"epoch_date_close_approach": 1441738200000,
	"relative_velocity": {
	"kilometers_per_second": "11.9557600601",
	"kilometers_per_hour": "43040.7362163935",
	"miles_per_hour": "26743.8396784585"
	},
	"miss_distance": {
	"astronomical": "0.1132399881",
	"lunar": "44.0503553709",
	"kilometers": "16940461.018585347",
	"miles": "10526314.3652659086"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3747356?api_key=DEMO_KEY"
	},
	"id": "3747356",
	"neo_reference_id": "3747356",
	"name": "(2016 EK158)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3747356",
	"absolute_magnitude_h": 20.56,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.2053784995,
	"estimated_diameter_max": 0.459240286
	},
	"meters": {
	"estimated_diameter_min": 205.3784995184,
	"estimated_diameter_max": 459.2402860401
	},
	"miles": {
	"estimated_diameter_min": 0.1276162436,
	"estimated_diameter_max": 0.2853585958
	},
	"feet": {
	"estimated_diameter_min": 673.8139963601,
	"estimated_diameter_max": 1506.6939000519
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-08",
	"close_approach_date_full": "2015-Sep-08 10:27",
	"epoch_date_close_approach": 1441708020000,
	"relative_velocity": {
	"kilometers_per_second": "16.957379392",
	"kilometers_per_hour": "61046.5658110207",
	"miles_per_hour": "37931.9619618534"
	},
	"miss_distance": {
	"astronomical": "0.2804751972",
	"lunar": "109.1048517108",
	"kilometers": "41958492.088949964",
	"miles": "26071798.0187349432"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3758838?api_key=DEMO_KEY"
	},
	"id": "3758838",
	"neo_reference_id": "3758838",
	"name": "(2016 RT)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3758838",
	"absolute_magnitude_h": 24.4,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.0350392641,
	"estimated_diameter_max": 0.0783501764
	},
	"meters": {
	"estimated_diameter_min": 35.0392641108,
	"estimated_diameter_max": 78.3501764334
	},
	"miles": {
	"estimated_diameter_min": 0.0217723826,
	"estimated_diameter_max": 0.0486845275
	},
	"feet": {
	"estimated_diameter_min": 114.9582192654,
	"estimated_diameter_max": 257.0543928497
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-08",
	"close_approach_date_full": "2015-Sep-08 01:02",
	"epoch_date_close_approach": 1441674120000,
	"relative_velocity": {
	"kilometers_per_second": "19.0983945697",
	"kilometers_per_hour": "68754.220451069",
	"miles_per_hour": "42721.1988130545"
	},
	"miss_distance": {
	"astronomical": "0.170705627",
	"lunar": "66.404488903",
	"kilometers": "25537198.19621449",
	"miles": "15868079.146520362"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/54191333?api_key=DEMO_KEY"
	},
	"id": "54191333",
	"neo_reference_id": "54191333",
	"name": "(2021 QP3)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=54191333",
	"absolute_magnitude_h": 22.737,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.0753624444,
	"estimated_diameter_max": 0.1685155487
	},
	"meters": {
	"estimated_diameter_min": 75.3624444246,
	"estimated_diameter_max": 168.515548684
	},
	"miles": {
	"estimated_diameter_min": 0.0468280375,
	"estimated_diameter_max": 0.104710675
	},
	"feet": {
	"estimated_diameter_min": 247.252122166,
	"estimated_diameter_max": 552.8725527443
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-08",
	"close_approach_date_full": "2015-Sep-08 00:22",
	"epoch_date_close_approach": 1441671720000,
	"relative_velocity": {
	"kilometers_per_second": "9.3106795473",
	"kilometers_per_hour": "33518.4463701775",
	"miles_per_hour": "20827.0590792917"
	},
	"miss_distance": {
	"astronomical": "0.3949696486",
	"lunar": "153.6431933054",
	"kilometers": "59086618.145208482",
	"miles": "36714722.0311497716"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	}
	],
	"2015-09-07": [
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/2440012?api_key=DEMO_KEY"
	},
	"id": "2440012",
	"neo_reference_id": "2440012",
	"name": "440012 (2002 LE27)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=2440012",
	"absolute_magnitude_h": 19.3,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.3669061375,
	"estimated_diameter_max": 0.8204270649
	},
	"meters": {
	"estimated_diameter_min": 366.9061375314,
	"estimated_diameter_max": 820.4270648822
	},
	"miles": {
	"estimated_diameter_min": 0.2279848336,
	"estimated_diameter_max": 0.5097895857
	},
	"feet": {
	"estimated_diameter_min": 1203.7603322587,
	"estimated_diameter_max": 2691.6899315481
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-07",
	"close_approach_date_full": "2015-Sep-07 07:32",
	"epoch_date_close_approach": 1441611120000,
	"relative_velocity": {
	"kilometers_per_second": "1.1630843052",
	"kilometers_per_hour": "4187.1034988155",
	"miles_per_hour": "2601.7032823612"
	},
	"miss_distance": {
	"astronomical": "0.4981690972",
	"lunar": "193.7877788108",
	"kilometers": "74525035.840942964",
	"miles": "46307709.9545183432"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3713989?api_key=DEMO_KEY"
	},
	"id": "3713989",
	"neo_reference_id": "3713989",
	"name": "(2015 FC35)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3713989",
	"absolute_magnitude_h": 22.1,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.1010543415,
	"estimated_diameter_max": 0.2259643771
	},
	"meters": {
	"estimated_diameter_min": 101.054341542,
	"estimated_diameter_max": 225.9643771094
	},
	"miles": {
	"estimated_diameter_min": 0.0627922373,
	"estimated_diameter_max": 0.140407711
	},
	"feet": {
	"estimated_diameter_min": 331.5431259047,
	"estimated_diameter_max": 741.3529669956
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-07",
	"close_approach_date_full": "2015-Sep-07 20:01",
	"epoch_date_close_approach": 1441656060000,
	"relative_velocity": {
	"kilometers_per_second": "8.763533811",
	"kilometers_per_hour": "31548.7217197058",
	"miles_per_hour": "19603.1487818916"
	},
	"miss_distance": {
	"astronomical": "0.3213750442",
	"lunar": "125.0148921938",
	"kilometers": "48077022.083475854",
	"miles": "29873676.2618966252"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3726788?api_key=DEMO_KEY"
	},
	"id": "3726788",
	"neo_reference_id": "3726788",
	"name": "(2015 RG2)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3726788",
	"absolute_magnitude_h": 26.7,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.0121494041,
	"estimated_diameter_max": 0.0271668934
	},
	"meters": {
	"estimated_diameter_min": 12.14940408,
	"estimated_diameter_max": 27.1668934089
	},
	"miles": {
	"estimated_diameter_min": 0.0075492874,
	"estimated_diameter_max": 0.0168807197
	},
	"feet": {
	"estimated_diameter_min": 39.8602508817,
	"estimated_diameter_max": 89.1302305717
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-07",
	"close_approach_date_full": "2015-Sep-07 17:58",
	"epoch_date_close_approach": 1441648680000,
	"relative_velocity": {
	"kilometers_per_second": "8.0871658927",
	"kilometers_per_hour": "29113.7972136669",
	"miles_per_hour": "18090.1813853476"
	},
	"miss_distance": {
	"astronomical": "0.0163786734",
	"lunar": "6.3713039526",
	"kilometers": "2450214.654065658",
	"miles": "1522492.7871077604"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3727036?api_key=DEMO_KEY"
	},
	"id": "3727036",
	"neo_reference_id": "3727036",
	"name": "(2015 RL35)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3727036",
	"absolute_magnitude_h": 26.3,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.0146067964,
	"estimated_diameter_max": 0.0326617897
	},
	"meters": {
	"estimated_diameter_min": 14.6067964271,
	"estimated_diameter_max": 32.6617897446
	},
	"miles": {
	"estimated_diameter_min": 0.0090762397,
	"estimated_diameter_max": 0.020295089
	},
	"feet": {
	"estimated_diameter_min": 47.92256199,
	"estimated_diameter_max": 107.1581062656
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-07",
	"close_approach_date_full": "2015-Sep-07 03:58",
	"epoch_date_close_approach": 1441598280000,
	"relative_velocity": {
	"kilometers_per_second": "3.5169616174",
	"kilometers_per_hour": "12661.0618226584",
	"miles_per_hour": "7867.0914419735"
	},
	"miss_distance": {
	"astronomical": "0.0692469329",
	"lunar": "26.9370568981",
	"kilometers": "10359193.665872923",
	"miles": "6436904.4607474174"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3727179?api_key=DEMO_KEY"
	},
	"id": "3727179",
	"neo_reference_id": "3727179",
	"name": "(2015 RH36)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3727179",
	"absolute_magnitude_h": 23.6,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.0506471459,
	"estimated_diameter_max": 0.1132504611
	},
	"meters": {
	"estimated_diameter_min": 50.6471458835,
	"estimated_diameter_max": 113.2504610618
	},
	"miles": {
	"estimated_diameter_min": 0.0314706677,
	"estimated_diameter_max": 0.0703705522
	},
	"feet": {
	"estimated_diameter_min": 166.1651821003,
	"estimated_diameter_max": 371.5566426699
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-07",
	"close_approach_date_full": "2015-Sep-07 11:50",
	"epoch_date_close_approach": 1441626600000,
	"relative_velocity": {
	"kilometers_per_second": "7.2717612888",
	"kilometers_per_hour": "26178.3406398515",
	"miles_per_hour": "16266.2028270233"
	},
	"miss_distance": {
	"astronomical": "0.1093379598",
	"lunar": "42.5324663622",
	"kilometers": "16356725.896225626",
	"miles": "10163598.1796045988"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3727662?api_key=DEMO_KEY"
	},
	"id": "3727662",
	"neo_reference_id": "3727662",
	"name": "(2015 RX83)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3727662",
	"absolute_magnitude_h": 22.9,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.0699125232,
	"estimated_diameter_max": 0.1563291544
	},
	"meters": {
	"estimated_diameter_min": 69.9125232246,
	"estimated_diameter_max": 156.3291544087
	},
	"miles": {
	"estimated_diameter_min": 0.0434416145,
	"estimated_diameter_max": 0.097138403
	},
	"feet": {
	"estimated_diameter_min": 229.3718026961,
	"estimated_diameter_max": 512.8909429502
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-07",
	"close_approach_date_full": "2015-Sep-07 21:51",
	"epoch_date_close_approach": 1441662660000,
	"relative_velocity": {
	"kilometers_per_second": "2.694557063",
	"kilometers_per_hour": "9700.4054267199",
	"miles_per_hour": "6027.4546941749"
	},
	"miss_distance": {
	"astronomical": "0.2895816212",
	"lunar": "112.6472506468",
	"kilometers": "43320793.722666844",
	"miles": "26918293.0014326872"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3727663?api_key=DEMO_KEY"
	},
	"id": "3727663",
	"neo_reference_id": "3727663",
	"name": "(2015 RY83)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3727663",
	"absolute_magnitude_h": 24.2,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.0384197891,
	"estimated_diameter_max": 0.0859092601
	},
	"meters": {
	"estimated_diameter_min": 38.4197891064,
	"estimated_diameter_max": 85.9092601232
	},
	"miles": {
	"estimated_diameter_min": 0.0238729428,
	"estimated_diameter_max": 0.0533815229
	},
	"feet": {
	"estimated_diameter_min": 126.0491808919,
	"estimated_diameter_max": 281.8545369825
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-07",
	"close_approach_date_full": "2015-Sep-07 16:55",
	"epoch_date_close_approach": 1441644900000,
	"relative_velocity": {
	"kilometers_per_second": "6.9802494143",
	"kilometers_per_hour": "25128.8978914704",
	"miles_per_hour": "15614.1199148417"
	},
	"miss_distance": {
	"astronomical": "0.0764899182",
	"lunar": "29.7545781798",
	"kilometers": "11442728.839194234",
	"miles": "7110181.9971550692"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3759353?api_key=DEMO_KEY"
	},
	"id": "3759353",
	"neo_reference_id": "3759353",
	"name": "(2016 RU33)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3759353",
	"absolute_magnitude_h": 27.5,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.008405334,
	"estimated_diameter_max": 0.0187948982
	},
	"meters": {
	"estimated_diameter_min": 8.4053340207,
	"estimated_diameter_max": 18.7948982439
	},
	"miles": {
	"estimated_diameter_min": 0.0052228308,
	"estimated_diameter_max": 0.0116786047
	},
	"feet": {
	"estimated_diameter_min": 27.5765560686,
	"estimated_diameter_max": 61.6630539546
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-07",
	"close_approach_date_full": "2015-Sep-07 16:34",
	"epoch_date_close_approach": 1441643640000,
	"relative_velocity": {
	"kilometers_per_second": "13.2144918467",
	"kilometers_per_hour": "47572.1706482885",
	"miles_per_hour": "29559.4968119879"
	},
	"miss_distance": {
	"astronomical": "0.2270491427",
	"lunar": "88.3221165103",
	"kilometers": "33966068.133246049",
	"miles": "21105536.0612875162"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3759690?api_key=DEMO_KEY"
	},
	"id": "3759690",
	"neo_reference_id": "3759690",
	"name": "(2016 RN41)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3759690",
	"absolute_magnitude_h": 31.028,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.0016555983,
	"estimated_diameter_max": 0.0037020304
	},
	"meters": {
	"estimated_diameter_min": 1.6555983184,
	"estimated_diameter_max": 3.7020303833
	},
	"miles": {
	"estimated_diameter_min": 0.0010287408,
	"estimated_diameter_max": 0.0023003343
	},
	"feet": {
	"estimated_diameter_min": 5.4317531869,
	"estimated_diameter_max": 12.1457693628
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-07",
	"close_approach_date_full": "2015-Sep-07 19:16",
	"epoch_date_close_approach": 1441653360000,
	"relative_velocity": {
	"kilometers_per_second": "13.4815494029",
	"kilometers_per_hour": "48533.5778505242",
	"miles_per_hour": "30156.8778593994"
	},
	"miss_distance": {
	"astronomical": "0.1205320685",
	"lunar": "46.8869746465",
	"kilometers": "18031340.714294095",
	"miles": "11204155.576264711"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3827337?api_key=DEMO_KEY"
	},
	"id": "3827337",
	"neo_reference_id": "3827337",
	"name": "(2018 RZ2)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3827337",
	"absolute_magnitude_h": 22.2,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.096506147,
	"estimated_diameter_max": 0.2157943048
	},
	"meters": {
	"estimated_diameter_min": 96.5061469579,
	"estimated_diameter_max": 215.7943048444
	},
	"miles": {
	"estimated_diameter_min": 0.059966121,
	"estimated_diameter_max": 0.134088323
	},
	"feet": {
	"estimated_diameter_min": 316.6212271853,
	"estimated_diameter_max": 707.9865871058
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-07",
	"close_approach_date_full": "2015-Sep-07 05:30",
	"epoch_date_close_approach": 1441603800000,
	"relative_velocity": {
	"kilometers_per_second": "18.513293253",
	"kilometers_per_hour": "66647.8557106394",
	"miles_per_hour": "41412.3856775359"
	},
	"miss_distance": {
	"astronomical": "0.4191352221",
	"lunar": "163.0436013969",
	"kilometers": "62701736.468136927",
	"miles": "38961052.3932945126"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3843641?api_key=DEMO_KEY"
	},
	"id": "3843641",
	"neo_reference_id": "3843641",
	"name": "(2019 QK4)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3843641",
	"absolute_magnitude_h": 20.8,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.1838886721,
	"estimated_diameter_max": 0.411187571
	},
	"meters": {
	"estimated_diameter_min": 183.8886720703,
	"estimated_diameter_max": 411.1875710413
	},
	"miles": {
	"estimated_diameter_min": 0.1142630881,
	"estimated_diameter_max": 0.2555000322
	},
	"feet": {
	"estimated_diameter_min": 603.309310875,
	"estimated_diameter_max": 1349.040630575
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-07",
	"close_approach_date_full": "2015-Sep-07 10:04",
	"epoch_date_close_approach": 1441620240000,
	"relative_velocity": {
	"kilometers_per_second": "38.349717322",
	"kilometers_per_hour": "138058.9823592299",
	"miles_per_hour": "85784.4826776004"
	},
	"miss_distance": {
	"astronomical": "0.3387069263",
	"lunar": "131.7569943307",
	"kilometers": "50669834.728726981",
	"miles": "31484775.3319990978"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/3986741?api_key=DEMO_KEY"
	},
	"id": "3986741",
	"neo_reference_id": "3986741",
	"name": "(2020 BY)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=3986741",
	"absolute_magnitude_h": 24.5,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.0334622374,
	"estimated_diameter_max": 0.0748238376
	},
	"meters": {
	"estimated_diameter_min": 33.4622374455,
	"estimated_diameter_max": 74.8238376074
	},
	"miles": {
	"estimated_diameter_min": 0.0207924639,
	"estimated_diameter_max": 0.0464933628
	},
	"feet": {
	"estimated_diameter_min": 109.7842471007,
	"estimated_diameter_max": 245.4850393757
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-07",
	"close_approach_date_full": "2015-Sep-07 05:39",
	"epoch_date_close_approach": 1441604340000,
	"relative_velocity": {
	"kilometers_per_second": "27.1899249786",
	"kilometers_per_hour": "97883.7299230601",
	"miles_per_hour": "60821.1431846925"
	},
	"miss_distance": {
	"astronomical": "0.4067666372",
	"lunar": "158.2322218708",
	"kilometers": "60851422.512182764",
	"miles": "37811320.6148355832"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	},
	{
	"links": {
	"self": "http://www.neowsapp.com/rest/v1/neo/54088823?api_key=DEMO_KEY"
	},
	"id": "54088823",
	"neo_reference_id": "54088823",
	"name": "(2020 WZ)",
	"nasa_jpl_url": "http://ssd.jpl.nasa.gov/sbdb.cgi?sstr=54088823",
	"absolute_magnitude_h": 26.9,
	"estimated_diameter": {
	"kilometers": {
	"estimated_diameter_min": 0.0110803882,
	"estimated_diameter_max": 0.0247765013
	},
	"meters": {
	"estimated_diameter_min": 11.0803882126,
	"estimated_diameter_max": 24.7765012606
	},
	"miles": {
	"estimated_diameter_min": 0.0068850319,
	"estimated_diameter_max": 0.0153953994
	},
	"feet": {
	"estimated_diameter_min": 36.3529808636,
	"estimated_diameter_max": 81.2877363957
	}
	},
	"is_potentially_hazardous_asteroid": false,
	"close_approach_data": [
	{
	"close_approach_date": "2015-09-07",
	"close_approach_date_full": "2015-Sep-07 19:56",
	"epoch_date_close_approach": 1441655760000,
	"relative_velocity": {
	"kilometers_per_second": "20.0777354148",
	"kilometers_per_hour": "72279.8474933128",
	"miles_per_hour": "44911.8863493865"
	},
	"miss_distance": {
	"astronomical": "0.462203352",
	"lunar": "179.797103928",
	"kilometers": "69144636.96606024",
	"miles": "42964485.121061712"
	},
	"orbiting_body": "Earth"
	}
	],
	"is_sentry_object": false
	}
	]
	}
	}`

func TestNeo(t *testing.T) {

	server := mockServer(200, mockData)
	defer server.Close()
	c := NewClient(WithBaseURL(server.URL))

	resp, err := c.NeoW()
	if err != nil {
		t.Errorf("returned error %s\n", err)
		return
	}

	correctResponse := NeoWResult{}

	err = json.Unmarshal([]byte(mockData), &correctResponse)
	if err != nil {
		t.Errorf("error when parsing the mock data: %s\n", err)
		return
	}

	if !reflect.DeepEqual(correctResponse, *resp) {
		t.Errorf("incorrect response from client: %+v\n", *resp)
	}
}

func TestNeoOptions(t *testing.T) {
	server := mockServer(200, mockData)
	defer server.Close()

	c := NewClient(WithBaseURL(server.URL))
	// start_date=2015-09-07&end_date=2015-09-08
	options := &NeoWOptions{
		StartDate: "2015-09-07",
		EndDate:   "2015-09-08",
	}
	resp, err := c.NeoWOpt(options)

	if err != nil {
		t.Errorf("unexpected error: %s\n", err)
		return
	}
	correctResponse := NeoWResult{}
	json.Unmarshal([]byte(mockData), &correctResponse)

	if !reflect.DeepEqual(correctResponse, *resp) {
		t.Errorf("incorrect response from client:\n%+v\n", *resp)
	}
}

func TestNeoMissingStartDate(t *testing.T) {
	options := &NeoWOptions{
		EndDate: "2015-09-08",
	}
	c := NewClient()

	_, err := c.NeoWOpt(options)

	if err == nil {
		t.Errorf("expected error received nil")
	}

}

func TestNeoWrongStartFormat(t *testing.T) {
	options := &NeoWOptions{
		StartDate: "March 20, 2021",
	}
	c := NewClient()

	_, err := c.NeoWOpt(options)

	if err == nil {
		t.Errorf("expected error received nil")
	}
}
func TestNeoWrongEndFormat(t *testing.T) {
	options := &NeoWOptions{
		StartDate: "2021-01-01",
		EndDate:   "March 20, 2021",
	}
	c := NewClient()

	_, err := c.NeoWOpt(options)

	if err == nil {
		t.Errorf("expected error received nil")
	}
}

func TestNeoUnexpectedError(t *testing.T) {
	c := NewClient(WithBaseURL("%"))

	_, err := c.NeoW()

	if err == nil {
		t.Errorf("expected error received nil")
	}
}

func TestNeoOptUnexpectedError(t *testing.T) {

	c := NewClient(WithBaseURL("%"))

	options := &NeoWOptions{
		StartDate: "2021-01-01",
		EndDate:   "2021-01-02",
	}
	_, err := c.NeoWOpt(options)

	if err == nil {
		t.Errorf("expected error received nil")
	}
}
