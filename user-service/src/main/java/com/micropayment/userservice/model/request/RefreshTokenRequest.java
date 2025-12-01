package com.micropayment.userservice.model.request;

import jakarta.validation.constraints.NotNull;
import lombok.Data;

/**
 *   Refresh Token Request.
 *   refreshToken         varchar(255) NOT NULL
 */
@Data
public class RefreshTokenRequest {
    @NotNull
    private String refreshToken;
}