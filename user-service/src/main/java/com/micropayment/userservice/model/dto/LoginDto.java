package com.micropayment.userservice.model.dto;

import lombok.Data;

/**
 * This is Login Dto that returns jwt token for user to use.
 */
@Data
public class LoginDto {
    private String token;
    private String refreshToken;
}
