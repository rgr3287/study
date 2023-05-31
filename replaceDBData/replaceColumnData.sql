DELIMITER // # 프로시저를 정의 할 경우에 문장 구분을 하기 위해 사용

CREATE PROCEDURE replace_string_proc(
    IN search_string VARCHAR(255),
    IN replace_string VARCHAR(255),
    IN db_name VARCHAR(255)
) # 프로시저 생성
BEGIN
    DECLARE done INT DEFAULT FALSE;
    DECLARE v_table_name VARCHAR(255); # 테이블
    DECLARE v_column_name VARCHAR(255); # 컬럼
    DECLARE cur CURSOR FOR # 커서로 만들 데이터 값들
SELECT table_name, column_name
FROM information_schema.columns
WHERE table_schema = db_name AND data_type IN ('char', 'varchar', 'text'); # 문자열 타입의 컬럼만 셀렉
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = TRUE; # 커서가 마지막에 도착할 때의 상태 값

    SET @query = ''; # 쿼리가 널 값일 경우 실행이 안되기 때문에 지정 해주 었습니다.

    OPEN cur; # 커서를 연다.

    read_loop: LOOP # loop가 돌아간다
        FETCH cur INTO v_table_name, v_column_name; # 커서로 만들어졌던 데이터를 돌린다.
        IF done THEN
            LEAVE read_loop;
END IF; # 커서가 마지막 로우일때 loop를 빠져오게 된다.

        SET @query = CONCAT(
                'UPDATE `', v_table_name, '`',
                ' SET `', v_column_name, '` = REPLACE(`', v_column_name, '`, ''', search_string, ''', ''', replace_string, ''') WHERE `', v_column_name, '` LIKE CONCAT(\'%\', ''', search_string, ''', \'%\');'
            ); # 각각의 테이블들의 문자열 컬럼에 교체하여 업데이트를 친다.

        PREPARE stmt FROM @query; # SQL 구문 선언
        EXECUTE stmt; # 미리 선언된 구문 실행 부분
        DEALLOCATE PREPARE stmt;  # 미리 선언된 구문 삭제
END LOOP;

CLOSE cur; # 커서 닫기

END //

DELIMITER ;

call replace_string_proc('SPURLINER', 'SNOWLINER', 'spice'); # 이전 문자, 바꿀 문자, 디비 이름

DROP PROCEDURE replace_string_proc; # 프로시저 삭제
