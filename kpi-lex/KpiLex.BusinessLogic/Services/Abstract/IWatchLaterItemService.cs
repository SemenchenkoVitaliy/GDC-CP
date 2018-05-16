using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using KpiLex.Domain.Models;

namespace KpiLex.BusinessLogic.Services.Abstract
{
    public interface IWatchLaterService : IService<WatchLaterItem, Guid>
    {
        Task<IReadOnlyCollection<WatchLaterItem>> GetWatchLaterItems(int studentId);
    }
}