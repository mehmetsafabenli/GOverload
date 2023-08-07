package client

import (
	"context"
	"net/url"

	"github.com/docker/docker/api/types"
) // RecordedStats contains both the container stats we've received from docker, and our own derived stats  from those container stats. When configuring a graph, you're basically specifying the path of a value in this struct
// CalculateContainerCPUPercentage calculates the cpu usage of the container as a percent of total CPU usage
// to calculate CPU usage, we take the increase in CPU time from the container since the last poll, divide that by the total increase in CPU time since the last poll, times by the number of cores, and times by 100 to get a percentage
// I'm not entirely sure why we need to multiply by the number of cores, but the numbers work

func (cli *Client) ContainerStats(ctx context.Context, containerID string, stream bool) (types.ContainerStats, error) {
	query := url.Values{}
	query.Set("stream", "0")
	if stream {
		query.Set("stream", "1")
	}

	resp, err := cli.get(ctx, "/containers/"+containerID+"/stats", query, nil)
	if err != nil {
		return types.ContainerStats{}, err
	}

	return types.ContainerStats{
		Body:   resp.body,
		OSType: getDockerOS(resp.header.Get("Server")),
	}, nil
}

// ContainerStatsOneShot gets a single stat entry from a container.
// It differs from `ContainerStats` in that the API should not wait to prime the stats
func (cli *Client) ContainerStatsOneShot(ctx context.Context, containerID string) (types.ContainerStats, error) {
	query := url.Values{}
	query.Set("stream", "0")
	query.Set("one-shot", "1")

	resp, err := cli.get(ctx, "/containers/"+containerID+"/stats", query, nil)
	if err != nil {
		return types.ContainerStats{}, err
	}

	return types.ContainerStats{
		Body:   resp.body,
		OSType: getDockerOS(resp.header.Get("Server")),
	}, nil
}
