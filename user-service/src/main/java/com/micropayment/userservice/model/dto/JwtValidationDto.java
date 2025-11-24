package com.micropayment.userservice.model.dto;


import lombok.Builder;
import lombok.Data;

import java.util.Map;

/**
 * This is response for jwt validation dto that returns needed value.
 */
@Data
@Builder
public class JwtValidationDto<T> {
    private boolean isValid;
    private T payload;
}
