create or replace view extended_employee as
select employee_id,
       first_name,
       second_name,
       employee.phone                          as employee_phone,
       COALESCE(company.company_id, -1)        as company_id,
       COALESCE(company.name, 'undefined')     as company_name,
       passport.passport_id                    as passport_id,
       passport.type                           as passport_type,
       passport.number                         as passport_number,
       COALESCE(department.department_id, -1)  as department_id,
       COALESCE(department.name, 'undefined')  as department_name,
       COALESCE(department.phone, 'undefined') as department_phone
from employee
         left join company using (company_id)
         join passport using (passport_id)
         left join department using (department_id);
