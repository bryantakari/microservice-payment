package com.micropayment.userservice.model.dto;

import lombok.Data;

/**
 * Jwt Token Dto class.
 */
@Data
public class JwtDto {
    private Integer userId;
    private String username;
    private String fullname;
}
