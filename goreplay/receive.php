<?php

    file_put_contents('./receive.txt',json_encode($_REQUEST),FILE_APPEND);

    echo json_encode($_REQUEST);

?>