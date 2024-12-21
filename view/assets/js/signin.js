$(document).ready(function () {
    $('#loginBtn').on('click', function (e) {
        e.preventDefault();
        const userId = $('#InputID').val();
        const password = $('#InputPW').val();

        if (!userId || !password) {
            alert('아이디와 비밀번호를 모두 입력해주세요!!');
            return;
        }

        $.ajax({
            url: '/auth/signin',
            method: 'POST',
            contentType: 'application/json; charset=utf-8',
            dataType: 'json',
            data: JSON.stringify({
                USER_ID: userId,
                PASSWORD: password,
            }),
            statusCode: {
                401: function () {
                    alert('아이디 또는 비밀번호가 올바르지 않습니다!!');
                },
                500: function () {
                    alert('서버 오류가 발생했습니다. 잠시 후 다시 시도해주세요!!');
                }
            }
        })
            .done(function (response) {
                console.log("서버 응답:", response);
                alert('로그인이 완료되었습니다.');
                window.location.href = '/';
            })
            .fail(function (jqXHR, textStatus, errorThrown) {
                console.error("AJAX 오류:", textStatus, errorThrown);
                alert('로그인 중 오류가 발생했습니다. 다시 시도해주세요.');
            });
    });
    // 자동 토큰 갱신 처리
    setInterval(function () {
        $.get('/auth/refresh-token')
            .fail(function (jqXHR) {
                if (jqXHR.status === 401) {
                    // 리프레시 토큰도 만료된 경우
                    window.location.href = '/login';
                }
            });
    }, 55 * 60 * 1000); // 55분마다 갱신

    $('#logoutBtn').on('click', function (e) {
        $.ajax({
            url: '/auth/signout',
            method: 'POST',
            success: function (response) {
                alert('로그아웃되었습니다');
                window.location.href = '/';
            },
            error: function (xhr, status, error) {
                alert('로그아웃 중 오류가 발생했습니다');
                console.error('로그아웃 오류:', error);
            }
        })
    });
});
