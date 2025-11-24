package com.micropayment.userservice.model.dto;

import lombok.Data;

/**
 * Register Dto class for register response.
 */
@Data
public class RegisterDto {
    private Integer userId;
    private String username;
}
