package com.micropayment.userservice.model.dto;

import lombok.Data;

/**
 * User Info Dto.
 */
@Data
public class UserInfoDto {
    private String username;
    private String email;
    private String fullname;
    private String lastLoginAt;
}
