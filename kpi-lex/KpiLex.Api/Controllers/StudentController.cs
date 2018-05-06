using System;
using KpiLex.Api.Constants;
using KpiLex.Api.Controllers.Abstract;
using KpiLex.Api.Mappers.Abstract;
using KpiLex.Api.Models.Request;
using KpiLex.Api.Models.Response;
using KpiLex.BusinessLogic.Services.Abstract;
using KpiLex.Domain.Models;
using Microsoft.AspNetCore.Mvc;

namespace KpiLex.Api.Controllers
{
    [Route(RouteConstants.StudentRoute)]
    public class StudentController : BaseRestApiController<StudentRequestModel, StudentResponseModel, Student, long>
    {
        public StudentController(IService<Student, long> service,
            IApiMapper<Student, StudentResponseModel> responseModelMapper,
            IApiMapper<StudentRequestModel, Student> requestModelMapper)
            : base(service, responseModelMapper, requestModelMapper)
        {
        }
    }
}