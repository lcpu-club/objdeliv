<?php

$sock = fsockopen("127.0.0.1:24032");

$write = fn ($data) => @fwrite($sock, $data);

$write("CONNECT /new-object?expire=10&id=3c271e2a-c9d6-4f9e-8d29-2a979b3c6407 HTTP/1.1\r\n\r\n");

for ($i = 0; $i < 1000; ++ $i) {
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    $write("123456783219382038192038901283908290839012839081290e22f9e2f9ed98eh\r\n");
    fflush($sock);
    usleep(200);
}

echo fgets($sock);

fclose($sock);
