package com.micropayment.userservice.service;

import com.micropayment.userservice.common.response.BaseValueResponse;
import com.micropayment.userservice.model.dto.LoginDto;
import com.micropayment.userservice.model.dto.RegisterDto;
import com.micropayment.userservice.model.request.LoginRequest;
import com.micropayment.userservice.model.request.RefreshTokenRequest;
import com.micropayment.userservice.model.request.RegisterRequest;

/**
 * Auth Service Interface.
 */
public interface AuthService {
    BaseValueResponse<RegisterDto> registerAccount(RegisterRequest request);

    BaseValueResponse<LoginDto> loginAccount(LoginRequest request);

    BaseValueResponse<LoginDto> refreshToken(RefreshTokenRequest request);
}
