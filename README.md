# Smart Way Task

## Описание
Простой веб-сервис для управления сотрудниками, написанный на языке Go. Сервис позволяет управлять сотрудниками, компаниями и отделами.

## Возможности
1. **Добавление сотрудников**: В ответ возвращается ID добавленного сотрудника.
2. **Удаление сотрудников по ID**.
3. **Вывод списка сотрудников для указанной компании**: Отображаются все доступные поля.
4. **Вывод списка сотрудников для указанного отдела компании**: Отображаются все доступные поля.
5. **Изменение сотрудника по его ID**: Изменяются только те поля, которые указаны в запросе.

## API Эндпоинты

### Добавление сотрудника
- **URL**: `api/v1/employees`
- **Метод**: `POST`
  - **Запрос**:
    ```json
    {
      "first_name": "Alexey",
      "second_name": "Ivanov",
      "employee_phone": "8-903-351-12-13",
      "passport":{
          "passport_type": "online updated",
          "passport_number": "passport_number "
      }
    }
  
### Обновление сотрудника
- **URL**: `api/v1/employees/{employee_id}`
- **Метод**: `PATCH`
- **Запрос**:
  ```json
  {
    "employee_company":{
        "company_id": 1
    },
    "department":{
        "department_id": 1
    }
  }
  
### Удаление сотрудника
- **URL**: `api/v1/employees/{employee_id}`
- **Метод**: `DELETE`

### Добавление компании
- **URL**: `api/v1/companies`
- **Метод**: `POST`
- **Запрос**:
  ```json
  {
    "company_name": "ao company_name",
  }

### Обновление компании
- **URL**: `api/v1/companies/{company_id}`
- **Метод**: `PATCH`
- **Запрос**:
  ```json
  {
    "company_name": "ao updated_company_name",
  }

### Вывод сотрудников компании
- **URL**: `api/v1/companies/{company_id}/employees`
- **Метод**: `GET`

### Добавление отдела
- **URL**: `api/v1/companies/{company_id}/departments`
- **Метод**: `POST`
- **Запрос**:
  ```json
  {
    "department_name": "financial dep",
    "department_phone": "+2 143 741 i6 11"
  }

### Обновление отдела
- **URL**: `api/v1/companies/{company_id}/departments/{department_id}`
- **Метод**: `PATCH`
- **Запрос**:
  ```json
  {
    "department_name": "updated_financial dep",
    "department_phone": "8 903 351 14 16"
  }

### Вывод сотрудников отдела
- **URL**: `api/v1/companies/{company_id}/departments/{department_id}/employees`
- **Метод**: `GET`