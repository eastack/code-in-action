#!/usr/bin/env python

import psycopg2
from openpyxl import Workbook

host = "localhost"
dbname = "volu"
user = "eastack"
password = "eastack"

conn_string = f"host={host} user={user} dbname={dbname} password={password}"
conn = psycopg2.connect(conn_string)
cursor = conn.cursor()

applications_data = []
missed_applications = []
with open('data/college-application.txt') as f:
    for app in f.readlines():
        application = app.split(":")
        school_code = application[0]
        major_code = application[1].strip()

        cursor.execute(
            f"select name, school_code, major_code, major_name, major_prob_explanation, lowest_score, lowest_rank from volu.public.major_admission_sd where school_code = '{school_code}' and major_code = '{major_code}'")

        rows = cursor.fetchall()

        if len(rows) > 1:
            print(f'query {school_code}:{major_code}')
            print(rows)

        if len(rows) == 0:
            missed_applications.append(app)

        for row in rows:
            applications_data.append(row)

    conn.commit()
    cursor.close()
    conn.close()

    print(missed_applications)

    wb = Workbook()
    ws = wb.active
    ws.title = "按分数排序"

    for num, app in enumerate(sorted(applications_data, key=lambda row: row[5]), start=1):
        for n, a in enumerate(app, start=1):
            ws.cell(row=num, column=n, value=a)

    wb.save('balances.xlsx')



