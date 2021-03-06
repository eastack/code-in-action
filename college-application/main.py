#!/usr/bin/env python

import psycopg2
import requests


def get_college_info(college_name):
    print(f'获取《{college_name}》信息:')
    host = "localhost"
    dbname = "volu"
    user = "eastack"
    password = "eastack"
    conn_string = f"host={host} user={user} dbname={dbname} password={password}"
    conn = psycopg2.connect(conn_string)
    print("Connection established")
    cursor = conn.cursor()

    college_admission = requests.get(f'https://quark.sm.cn/api/rest?method=gaokaochoice.getCollegeAdmission&format=json&uc_param_str=dnntnwvepffrgibijbprsvdsdicheiniutst&dn=58858880501-ff776583&nt=5&nw=0&ve=5.1.2.182&pf=3300&fr=android&gi=bTkwBPDkznbaL07h0FRT8GlRCyhE&bi=36734&pr=ucpro&sv=released&ds=AAMf9zUgmue1AAaI51b8iE7E3oygSoopd5dbyP8wX5eWkA%3D%3D&de=AAPFEjJOSzg%252BPmafmMA5BepwvIjQK0pdlj%252BdnzpLEs9%252BQg%253D%253D&ch=kk%40other_zcwyy6&ni=bTkwBDwzd5QwrNPPAYvQhaoxtfgVcF92EUQG5ir%2Bk7W65d0%3D&ut=AAMf9zUgmue1AAaI51b8iE7E3oygSoopd5dbyP8wX5eWkA%3D%3D&st=st83a621a92hs4222wxp4xg5mni2wx2w&location=%E5%B1%B1%E4%B8%9C&aos=%E7%BB%BC%E5%90%88&score=499&subjects=%E7%89%A9%E7%90%86%2F%E7%94%9F%E7%89%A9%2F%E6%80%9D%E6%83%B3%E6%94%BF%E6%B2%BB&batch=%E6%99%AE%E9%80%9A%E7%B1%BB%E4%B8%80%E6%AE%B5&major=&rank=153262&priority=college&college_name={college_name}&tab=1&kpsWg=AATF9KMBIlhosG4hAqpX12%2B6vq1IezK7dFPhYQhhdE%2FhcBUvdTed%2FhunRaMdFEXHz4o6DEAmEC9XC3wkqVeop6C0%2BpVi2ziSKjZAe4e49o%2F8gQ%3D%3D&kps=AATF9KMBIlhosG4hAqpX12%2B6vq1IezK7dFPhYQhhdE%2FhcBUvdTed%2FhunRaMdFEXHz4o6DEAmEC9XC3wkqVeop6C0%2BpVi2ziSKjZAe4e49o%2F8gQ%3D%3D&utdid=AAMf9zUgmue1AAaI51b8iE7E3oygSoopd5dbyP8wX5eWkA%3D%3D&st=st83a621a92hs4222wxp4xg5mni2wx2w&timeStamp=1625319810236').json()['data']
    if college_admission['status_code'] == 301:
        print('失败')
        return
    else:
        print('成功')

    for major_admission_info in college_admission['major_admissions_info'][0]['majors']:
        if major_admission_info['major_prob_explanation'] != '难':
            major_name = major_admission_info['major_name']
            major_id = major_admission_info['major_id']
            major_admission = requests.get(
                f'https://quark.sm.cn/api/rest?method=gaokaochoice.getMajorAdmission&format=json&uc_param_str=dnntnwvepffrgibijbprsvdsdicheiniutst&dn=58858880501-ff776583&nt=5&nw=0&ve=5.1.2.182&pf=3300&fr=android&gi=bTkwBPDkznbaL07h0FRT8GlRCyhE&bi=36734&pr=ucpro&sv=released&ds=AAMf9zUgmue1AAaI51b8iE7E3oygSoopd5dbyP8wX5eWkA%3D%3D&de=AAPFEjJOSzg%252BPmafmMA5BepwvIjQK0pdlj%252BdnzpLEs9%252BQg%253D%253D&ch=kk%40other_zcwyy6&ni=bTkwBDwzd5QwrNPPAYvQhaoxtfgVcF92EUQG5ir%2Bk7W65d0%3D&ut=AAMf9zUgmue1AAaI51b8iE7E3oygSoopd5dbyP8wX5eWkA%3D%3D&st=st83a621a92hs4222wxp4xg5mni2wx2w&location=%E5%B1%B1%E4%B8%9C&aos=%E7%BB%BC%E5%90%88&score=499&subjects=%E7%89%A9%E7%90%86%2F%E7%94%9F%E7%89%A9%2F%E6%80%9D%E6%83%B3%E6%94%BF%E6%B2%BB&batch=%E6%99%AE%E9%80%9A%E7%B1%BB%E4%B8%80%E6%AE%B5&major=&rank=153262&priority=college&college_name={college_name}&tab=1&major_name={major_name}&major_id={major_id}&kpsWg=AATF9KMBIlhosG4hAqpX12%2B6vq1IezK7dFPhYQhhdE%2FhcBUvdTed%2FhunRaMdFEXHz4o6DEAmEC9XC3wkqVeop6C0%2BpVi2ziSKjZAe4e49o%2F8gQ%3D%3D&kps=AATF9KMBIlhosG4hAqpX12%2B6vq1IezK7dFPhYQhhdE%2FhcBUvdTed%2FhunRaMdFEXHz4o6DEAmEC9XC3wkqVeop6C0%2BpVi2ziSKjZAe4e49o%2F8gQ%3D%3D&utdid=AAMf9zUgmue1AAaI51b8iE7E3oygSoopd5dbyP8wX5eWkA%3D%3D&st=st83a621a92hs4222wxp4xg5mni2wx2w&timeStamp=1625327884759').json()[
                'data']
            for major_admissions_history_item in major_admission['history_info']:
                if major_admissions_history_item['year'] == '2020':
                    if major_admission['major_admissions_info']['school_fee'] == '':
                        school_fee = -1
                    else:
                        school_fee = major_admission['major_admissions_info']['school_fee']

                    cursor.execute(
                        "insert into volu.public.major_admission_sd" +
                        "("
                        "name, "
                        "major_code, "
                        "major_name, "
                        "lowest_score, "
                        "lowest_rank, "
                        "major_remark, "
                        "plan_num, "
                        "batch, "
                        "school_nature, "
                        "school_code, "
                        "location, "
                        "school_fee, "
                        "abnormal_remark, "
                        "major_prob_explanation) values " +
                        f"('{college_admission['name']}', "
                        f"'{major_admission_info['major_code']}', "
                        f"'{major_admission_info['major_name']}', "
                        f"'{major_admissions_history_item['low_score']}', "
                        f"'{major_admissions_history_item['low_rank']}', "
                        f"'{major_admissions_history_item['major_remark']}', "
                        f"'{major_admissions_history_item['plan_num']}', "
                        f"'{major_admissions_history_item['batch']}', "
                        f"'{college_admission['school_nature']}', "
                        f"'{college_admission['school_code']}', "
                        f"'{college_admission['location']}', "
                        f"'{school_fee}', "
                        f"'{college_admission['abnormal_remark']}', "
                        f"'{major_admission_info['major_prob_explanation']}')")

    conn.commit()
    cursor.close()
    conn.close()


with open('data/sd.txt') as f:
    colleges = f.readlines()
    for college in colleges:
        get_college_info(college)
