using Grpc.Net.Client;
using Microsoft.Extensions.Caching.Distributed;

namespace ProtoCache;

public class ProtoDistributedCache : IDistributedCache
{
    public ProtoDistributedCache(
        string address
    )
    {
        if (string.IsNullOrEmpty(address)) throw new ArgumentException($"'{nameof(address)}' cannot be null or empty.", nameof(address));

        Address = address;
    }

    private readonly string Address;

    public byte[]? Get(string key)
    {
        using var channel = GrpcChannel.ForAddress(Address);
        var client = new protoCache.protoCacheClient(channel);
        var request = new GetCacheItemRequest
        {
            Key = key,
        };
        var response = client.GetCacheItem(request);
        return response?.Value.ToByteArray();
    }

    public async Task<byte[]?> GetAsync(string key, CancellationToken token = default)
    {
        using var channel = GrpcChannel.ForAddress(Address);
        var client = new protoCache.protoCacheClient(channel);
        var request = new GetCacheItemRequest
        {
            Key = key,
        };
        var response = await client.GetCacheItemAsync(request, cancellationToken: token);
        return response?.Value.ToByteArray();
    }

    public void Refresh(string key)
    {
        throw new NotImplementedException();
    }

    public Task RefreshAsync(string key, CancellationToken token = default)
    {
        throw new NotImplementedException();
    }

    public void Remove(string key)
    {
        using var channel = GrpcChannel.ForAddress(Address);
        var client = new protoCache.protoCacheClient(channel);
        var request = new RemoveCacheItemRequest
        {
            Key = key,
        };
        client.RemoveCacheItem(request);
    }

    public async Task RemoveAsync(string key, CancellationToken token = default)
    {
        using var channel = GrpcChannel.ForAddress(Address);
        var client = new protoCache.protoCacheClient(channel);
        var request = new RemoveCacheItemRequest
        {
            Key = key,
        };
        await client.RemoveCacheItemAsync(request, cancellationToken: token);
    }

    public void Set(string key, byte[] value, DistributedCacheEntryOptions options)
    {
        using var channel = GrpcChannel.ForAddress(Address);
        var client = new protoCache.protoCacheClient(channel);
        var request = new SetCacheItemRequest
        {
            Key = key,
            Value = Google.Protobuf.ByteString.CopyFrom(value)
        };
        client.SetCacheItem(request);
    }

    public async Task SetAsync(string key, byte[] value, DistributedCacheEntryOptions options, CancellationToken token = default)
    {
        using var channel = GrpcChannel.ForAddress(Address);
        var client = new protoCache.protoCacheClient(channel);
        var request = new SetCacheItemRequest
        {
            Key = key,
            Value = Google.Protobuf.ByteString.CopyFrom(value)
        };
        await client.SetCacheItemAsync(request, cancellationToken: token);
    }
}