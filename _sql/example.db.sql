-- Небольшой набор тестовых данных для CarRental
UPDATE "RENTAL" SET "DUMMY"=1;
DELETE FROM "CARRENT";
DELETE FROM "RENTAL";
DELETE  FROM "CARGOODS";
DELETE FROM "CAR";
DELETE FROM "DEPARTMENT";
DELETE FROM "MODEL";
DELETE FROM "CARTYPE";

--TYPES
INSERT INTO "CARTYPE" ("NAME","CODE") VALUES ('Легковой автомобиль','AUTO');
INSERT INTO "CARTYPE" ("NAME","CODE") VALUES ('Мотоцикл','MOTORBIKE');
INSERT INTO "CARTYPE" ("NAME","CODE") VALUES ('Мопед','MOPED');
--MODEL CARS
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW X5',
         (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Corolla FX',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Corolla II',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Corolla Levin',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Corolla Rumion',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Corolla Runx',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Corolla Spacio',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Corolla Van',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Corolla Verso',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Corolla Wagon',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Corona',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Corona Exiv',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Corona Premio',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Corona SF',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Corona Wagon',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Corsa',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Cressida',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Cresta',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Crown',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW 1 (F20)',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW 2',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW F45',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW 3',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW F30',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW 4',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW M4',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW 5',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW 5 Series Gran Turismo',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW 6',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW 6 (F13)',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW 7',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW 8',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW 320si WTCC',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW 328',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW 501/502',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW 507',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW 700',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW 3200 CS',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='AUTO' ))
--MODEL MOTORBIKES
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Honda CRF250F/450F',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Honda XR250/400/650',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Yamaha TT250R Open Enduro',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Suzuki DR-Z400',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('KTM EX-C250',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Honda CRF250R',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Yamaha YZ250',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Suzuki RM250(2T)/RM-Z250(4T)',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Урал Т',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Урал GEAR-UP',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Днепр 11',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Днепр 16',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('K750',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Suzuki GSX-R 1000',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Honda CBR 1000RR',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Yamaha YZF-R1',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Honda CBR600RR',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Kawasaki Ninja ZX-6RR',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Triumph Daytona 675',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Yamaha YZF-R6           ',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Harley-Davidson V-Rod',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Yamaha V-Max',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Honda X4',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Kawasaki Eliminator',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BMW G650 X',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Suzuki DR-Z 400 SM',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Kawasaki D-Tracker 250',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Suzuki 250SB',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('BM Motard 200',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('KTM 690 SMC',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Husqvarna SMR 610',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Honda XR400 SM',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Honda FMX 650',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Yamaha XT660X',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Aprilia SXV 4.5-5.5',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOTORBIKE' ));
--MODEL MOPED

INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Гауя',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Рига-5',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Рига-7',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Рига-11',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Рига-13',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Кроха',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Рига-1',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Рига-3',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Рига-4',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Верховина-3',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Верховина-4',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Верховина-5',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Верховина-6',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Рига-16',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Рига-22',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Дельта',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Карпаты',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Верховина-7',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Иж 2.673 Корнет',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('ЗИД-50',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Птаха',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Simson S51',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Ryś MR1',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Komar',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Romet',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Jawa 50 typ 20',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Berva',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Carpati',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Mobra',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));
INSERT INTO "MODEL"("NAME", "CTYPE")
  VALUES('Балкан-50',
           (SELECT "ID" FROM "CARTYPE"  WHERE "CODE"='MOPED' ));


INSERT INTO "DEPARTMENT"("NAME") VALUES ('Амстердам');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Антверпен');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Афины');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Барселона');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Берлин');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Брюгге');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Варшава');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Вашингтон');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Вена');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Венеция');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Гагра');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Гамбург');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Геленджик');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Дрезден');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Дубай');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Дубровник');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Евпатория');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Ейск');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Екатеринбург');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Женева');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Загреб');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Иерусалим');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Ижевск');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Иркутск');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Йошкар-Ола');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Казань');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Кёльн');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Киев');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Лас-Вегас');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Лондон');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Лос-Анджелес');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Мадрид');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Милан');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Минск');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Москва');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Мюнхен');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Неаполь');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Нижний Новгород');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Новосибирск');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Нью-Йорк');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Одесса');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Омск');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Осло');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Париж');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Пекин');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Прага');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Рим');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Рио-де-Жанейро');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Ростов-на-Дону');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Санкт-Петербург');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Севастополь');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('София');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Сочи');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Таллин');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Тбилиси');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Тель-Авив');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Ульяновск');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Уфа');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Филадельфия');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Флоренция');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Хайфа');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Ханой');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Харьков');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Хельсинки');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Хошимин');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Цюрих');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Чебоксары');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Челябинск');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Чикаго');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Шанхай');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Штутгарт');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Щецин');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Эйлат');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Эйндховен');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Юрмала');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Якутск');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Ялта');
INSERT INTO "DEPARTMENT"("NAME") VALUES ('Ярославль');

-- ГЕНЕРАЦИЯ ТЕСТОВЫХ ДАННЫХ ДЛЯ CAR+CARGOODS
do $$
DECLARE R RECORD;
DECLARE NUM INT;
DECLARE CNEW INT; -- Количество зкземпляров данной модели
DECLARE MAXCARMODEL FLOAT=5.00;
declare DEPTS integer[];
DECLARE RNDDEP integer;
DECLARE IDCAR integer;
begin
  DEPTS = ARRAY(SELECT "ID" FROM "DEPARTMENT");
  NUM = 1;
  FOR r in SELECT "ID" FROM "MODEL"
  LOOP
    CNEW = random()*MAXCARMODEL+1;
    LOOP
      IF CNEW=0 THEN
        EXIT;
      END IF;
      NUM = NUM + 1;
      CNEW = CNEW - 1;
      INSERT INTO "CAR"("MODEL","REGNUM") VALUES (R."ID",'X'||NUM||'XX77RUS') RETURNING "CAR"."ID" INTO IDCAR;
      LOOP
        RNDDEP = random() * (array_length(DEPTS,1)+1);
        IF DEPTS[RNDDEP] IS NOT NULL THEN
          EXIT;
        END IF;
      END LOOP;
      --RAISE NOTICE '%, %', array_length(DEPTS,1), RNDDEP;
      INSERT INTO "CARGOODS"( "DEPT","CAR") VALUES( DEPTS[RNDDEP],IDCAR );
    END LOOP;
  END LOOP;
end $$;