using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using KpiLex.Api.Constants;
using KpiLex.Api.Mappers.Abstract;
using KpiLex.Api.Models;
using KpiLex.Api.Models.Request;
using KpiLex.Api.Models.Response;
using KpiLex.BusinessLogic.Services.Abstract;
using KpiLex.Domain.Models;
using Microsoft.AspNetCore.Mvc;

namespace KpiLex.Api.Controllers
{
    [Route(RouteConstants.CourseCommentRoute)]
    public class CourseCommentController : Controller
    {
        private readonly ICourseCommentService _courseCommentService;
        private readonly IApiMapper<CourseComment, CourseCommentResponseModel> _responseModelMapper;
        private readonly IApiMapper<CourseCommentRequestModel, CourseComment> _requestModelMapper;

        public CourseCommentController(
            ICourseCommentService courseCommentService,
            IApiMapper<CourseComment, CourseCommentResponseModel> responseModelMapper,
            IApiMapper<CourseCommentRequestModel, CourseComment> requestModelMapper)
        {
            _courseCommentService = courseCommentService;
            _responseModelMapper = responseModelMapper;
            _requestModelMapper = requestModelMapper;
        }
        
        [HttpGet]
        public async Task<IReadOnlyCollection<CourseCommentResponseModel>> GetCourseComments([FromRoute] int courseId)
        {
            var courseComments = await _courseCommentService.GetCourseComments(courseId);
            var responseModels = courseComments.Select(c => _responseModelMapper.Map(c)).ToList();
            return responseModels;
        }
        
        [HttpPost]
        public async Task<CourseCommentResponseModel> AddCourseComment(
            [FromRoute] int courseId,
            [FromBody] CourseCommentRequestModel courseCommentRequestModel)
        {
            var model = _requestModelMapper.Map(courseCommentRequestModel);
            model = await _courseCommentService.AddAsync(model);
            return _responseModelMapper.Map(model);
        }
        
        [HttpPut]
        public async Task<IActionResult> UpdateCourseComment(
            [FromRoute] int courseId,
            [FromBody] CourseCommentRequestModel courseCommentRequestModel)
        {
            var model = _requestModelMapper.Map(courseCommentRequestModel);
            await _courseCommentService.UpdateAsync(model);
            return Ok();
        }
        
        [HttpDelete, Route("{courseCommentId}")]
        public async Task<IActionResult> DeleteCourseComment(
            [FromRoute] int courseId,
            [FromRoute] int courseCommentId)
        {
            await _courseCommentService.RemoveAsync(courseCommentId);
            return Ok();
        }
    }
}