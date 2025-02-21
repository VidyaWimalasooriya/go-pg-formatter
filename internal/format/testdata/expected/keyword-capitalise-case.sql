-- Test
Select
    a,
    b,
    c
From
    tablea
    Join tableb On (tablea.a = tableb.a)
    Join tablec On (tablec.a = tableb.a)
    Left Outer Join tabled On (tabled.a = tableb.a)
    Left Join tablee On (tabled.a = tableb.a)
Where
    tablea.x = 1
    And tableb.y = 1
Group By
    tablea.a,
    tablec.c
Order By
    tablea.a,
    tablec.c;

