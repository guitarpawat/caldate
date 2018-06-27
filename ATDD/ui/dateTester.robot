*** Settings ***
Library    SeleniumLibrary

*** Test Cases ***
Case 1
    Open Browser    http://localhost:8000    gc
    InputStartDate    4    1    2018
    InputEndDate    4    7    2018
    Click Button    button
    CheckFromTo    Thursday, 4 January 2018    Wednesday, 4 July 2018
    CheckDateYear    182 days    6 months, 1 day
    CheckTime    15,724,800 seconds    262,080 minutes    4,368 hours
    CheckWeekAndPercent    26 weeks    49.86% of 2018
    Close Browser

Case 2
    Open Browser    http://localhost:8000    gc
    InputStartDate    11    4    1977
    InputEndDate    27    7    2018
    Click Button    button
    CheckFromTo    Monday, 11 April 1977    Friday, 27 July 2018
    CheckDateYear    15,083 days    41 years, 3 months, 17 days
    CheckTime    1,303,171,200 seconds    21,719,520 minutes    361,992 hours
    CheckWeekAndPercent    2154 weeks and 5 days    4132.33% of a common year (365 days)
    Close Browser

Case 3
    Open Browser    http://localhost:8000    gc
    InputStartDate    4    1    1997
    InputEndDate    26    6    2018
    Click Button    button
    CheckFromTo    Saturday, 4 January 1997    Tuesday, 26 June 2018
    CheckDateYear    7,844 days    21 years, 5 months, 23 days
    CheckTime    677,721,600 seconds    11,295,360 minutes    188,256 hours
    CheckWeekAndPercent    1120 weeks and 4 days    2149.04% of a common year (365 days)
    Close Browser

Case 4
    Open Browser    http://localhost:8000    gc
    InputStartDate    3    1    1997
    InputEndDate    26    6    2018
    Click Button    button
    CheckFromTo    Friday, 3 January 1997    Tuesday, 26 June 2018
    CheckDateYear    7,845 days    21 years, 5 months, 24 days
    CheckTime    677,808,000 seconds    11,296,800 minutes    188,280 hours
    CheckWeekAndPercent    1120 weeks and 5 days    2149.32% of a common year (365 days)
    Close Browser

Case 5
    Open Browser    http://localhost:8000    gc
    InputStartDate    23    8    1998
    InputEndDate    26    6    2018
    Click Button    button
    CheckFromTo    Sunday, 23 August 1998    Tuesday, 26 June 2018
    CheckDateYear    7,248 days    19 years, 10 months, 4 days
    CheckTime    626,227,200 seconds    10,437,120 minutes    173,952 hours
    CheckWeekAndPercent    1035 weeks and 3 days    1985.75% of a common year (365 days)
    Close Browser

*** Keywords ***
InputStartDate
    [Arguments]    ${StartDate}    ${StartMonth}    ${StartYear}
    Input Text    StartDate    ${StartDate}
    Input Text    StartMonth    ${StartMonth}
    Input Text    StartYear    ${StartYear}

InputEndDate
    [Arguments]    ${EndDate}    ${EndMonth}    ${EndYear}
    Input Text    EndDate    ${EndDate}
    Input Text    EndMonth    ${EndMonth}
    Input Text    EndYear    ${EndYear}

CheckFromTo
    [Arguments]    ${FromDate}    ${ToDate}
    Element Text Should Be    from    ${FromDate}
    Element Text Should Be    to    ${ToDate}

CheckDateYear
    [Arguments]    ${CheckDate}    ${CheckYear}
    Element Text Should Be    date    ${CheckDate}
    Element Text Should Be    years    ${CheckYear}

CheckTime
    [Arguments]    ${CheckSecond}    ${CheckMinute}    ${CheckHour}
    Element Text Should Be    seconds     ${CheckSecond}
    Element Text Should Be    minutes    ${CheckMinute}
    Element Text Should Be    hours    ${CheckHour}
    
CheckWeekAndPercent
    [Arguments]    ${CheckWeek}    ${CheckPercent}
    Element Text Should Be    weeks    ${CheckWeek}
    Element Text Should Be    percent    ${CheckPercent}