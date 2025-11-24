package com.micropayment.userservice.mapper;

import com.micropayment.userservice.model.dto.JwtDto;
import com.micropayment.userservice.model.dto.RegisterDto;
import com.micropayment.userservice.model.dto.UserInfoDto;
import com.micropayment.userservice.model.entity.Account;
import com.micropayment.userservice.model.request.RegisterRequest;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.MappingConstants;
import org.mapstruct.ReportingPolicy;

@Mapper(unmappedTargetPolicy = ReportingPolicy.IGNORE,
        componentModel = MappingConstants.ComponentModel.SPRING)
public interface UserMapper {


    Account mappingRegisterToAccount(RegisterRequest request);
    @Mapping(source = "id",target = "userId")
    RegisterDto mappingAccountToDto(Account account);
    @Mapping(source = "id",target = "userId")
    JwtDto mappingAccountToJwtDto(Account account);

    UserInfoDto mappingToUserInfoDto(Account account);
}
