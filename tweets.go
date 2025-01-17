package twitter

import (
	"fmt"
	"net/url"
	"time"
)

// GetUserMentions returns Tweets mentioning a single user specified by the requested user ID.
// Endpoint URL: https://api.twitter.com/2/users/:id/mentions
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-mentions
// Authentication Methods: OAuth 1.0a User Context, OAuth 2.0 Bearer Token
// Rate Limit: 450/15m (app), 180/15m (user)
func (api *Twitter) GetUserMentions(id string, v url.Values, options ...QueueOption) (chan *Data, chan error) {
	// create the queue to process requests
	queue := NewQueue(15*time.Minute/1500, 15*time.Minute, true, make(chan *Request), make(chan *Response), options...)
	// create the temp results channel
	data := make(chan *Data)
	errors := make(chan error)
	// create the request object
	request, _ := NewRquest("GET", fmt.Sprintf("%s/users/%s/mentions", api.baseURL, id), v, nil)
	// start the requests channel processor
	go queue.processRequests(api)
	// add the 1st request to the channel
	queue.requestsChannel <- request

	// async process the response channel
	go (func(q *Queue, d chan *Data, e chan error, req *Request) {
		// on done close channels
		// close data channel
		defer close(d)
		// close error channel
		defer close(e)

		// listen channel
		for {
			// capture the response and channel state
			res, ok := <-q.responseChannel
			// break the loop if the channel is closed
			if !ok {
				break
			}

			// send the results to the data channel
			d <- &res.Results
			// send errors to error channel
			if res.Error != nil {
				e <- res.Error
			}
			// if there is a next page, transform the original request object
			// by setting the `pagination_token` parameter to get the next page
			if res.Results.Meta != nil && res.Results.Meta.NextToken != "" && q.auto {
				// create new url values and add the pagination token
				nv := url.Values{}
				nv.Add("pagination_token", res.Results.Meta.NextToken)

				// update request's url Values
				req.UpdateURLValues(nv)
				// reset request's results
				req.ResetResults()

				// add next request to the channel
				q.requestsChannel <- req

				//go to start
				continue
			}
			// we are done! break the loop and close the channels
			break
		}
	})(queue, data, errors, request)

	// return the data channel
	return data, errors
}

// GetUserTweets returns Tweets composed by a single user, specified by the requested user ID.
// Endpoint URL: https://api.twitter.com/2/users/:id/tweets
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/timelines/api-reference/get-users-id-tweets
// Authentication Methods: OAuth 1.0a User Context, OAuth 2.0 Bearer Token
// Rate Limit: 1500/15m (app), 900/15m (user)
func (api *Twitter) GetUserTweets(id string, v url.Values, options ...QueueOption) (chan *Data, chan error) {
	// create the queue to process requests
	queue := NewQueue(15*time.Minute/1500, 15*time.Minute, true, make(chan *Request), make(chan *Response), options...)
	// create the temp results channel
	data := make(chan *Data)
	errors := make(chan error)
	// create the request object
	request, _ := NewRquest("GET", fmt.Sprintf("%s/users/%s/tweets", api.baseURL, id), v, nil)
	// start the requests channel processor
	go queue.processRequests(api)
	// add the 1st request to the channel
	queue.requestsChannel <- request

	// async process the response channel
	go (func(q *Queue, d chan *Data, e chan error, req *Request) {
		// on done close channels
		// close data channel
		defer close(d)
		// close error channel
		defer close(e)

		// listen channel
		for {
			// capture the response and channel state
			res, ok := <-q.responseChannel
			// break the loop if the channel is closed
			if !ok {
				break
			}

			// send the results to the data channel
			d <- &res.Results
			// send errors to error channel
			if res.Error != nil {
				e <- res.Error
			}
			// if there is a next page, transform the original request object
			// by setting the `pagination_token` parameter to get the next page
			if res.Results.Meta != nil && res.Results.Meta.NextToken != "" && q.auto {
				// create new url values and add the pagination token
				nv := url.Values{}
				nv.Add("pagination_token", res.Results.Meta.NextToken)

				// update request's url Values
				req.UpdateURLValues(nv)
				// reset request's results
				req.ResetResults()

				// add next request to the channel
				q.requestsChannel <- req

				//go to start
				continue
			}
			// we are done! break the loop and close the channels
			break
		}
	})(queue, data, errors, request)

	// return the data channel
	return data, errors
}

// GetTweets returns a variety of information about the Tweet specified by the requested ID or list of IDs.
// Endpoint URL: https://api.twitter.com/2/tweets
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
// Authentication Methods: OAuth 1.0a User Context, OAuth 2.0 Bearer Token
// Rate Limit: 300/15m (app), 900/15m (user)
func (api *Twitter) GetTweets(v url.Values, options ...QueueOption) (chan *Data, chan error) {
	// create the queue to process requests
	queue := NewQueue(15*time.Minute/1500, 15*time.Minute, true, make(chan *Request), make(chan *Response), options...)
	// create the temp results channel
	data := make(chan *Data)
	errors := make(chan error)
	// create the request object
	request, _ := NewRquest("GET", fmt.Sprintf("%s/tweets", api.baseURL), v, nil)
	// start the requests channel processor
	go queue.processRequests(api)
	// add the 1st request to the channel
	queue.requestsChannel <- request

	// async process the response channel
	go (func(q *Queue, d chan *Data, e chan error, req *Request) {
		// on done close channels
		// close data channel
		defer close(d)
		// close error channel
		defer close(e)

		// listen channel
		for {
			// capture the response and channel state
			res, ok := <-q.responseChannel
			// break the loop if the channel is closed
			if !ok {
				break
			}

			// send the results to the data channel
			d <- &res.Results
			// send errors to error channel
			if res.Error != nil {
				e <- res.Error
			}
			// if there is a next page, transform the original request object
			// by setting the `pagination_token` parameter to get the next page
			if res.Results.Meta != nil && res.Results.Meta.NextToken != "" && q.auto {
				// create new url values and add the pagination token
				nv := url.Values{}
				nv.Add("pagination_token", res.Results.Meta.NextToken)

				// update request's url Values
				req.UpdateURLValues(nv)
				// reset request's results
				req.ResetResults()

				// add next request to the channel
				q.requestsChannel <- req

				//go to start
				continue
			}
			// we are done! break the loop and close the channels
			break
		}
	})(queue, data, errors, request)

	// return the data channel
	return data, errors
}

// GetTweetByID returns a variety of information about a single Tweet specified by the requested ID.
// Endpoint URL: https://api.twitter.com/2/tweets/:id
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference//get-tweets-id
// Authentication Methods: OAuth 1.0a User Context, OAuth 2.0 Bearer Token
// Rate Limit: 300/15m (app), 900/15m (user)
func (api *Twitter) GetTweetByID(id string, v url.Values, options ...QueueOption) (chan *Data, chan error) {
	// create the queue to process requests
	queue := NewQueue(15*time.Minute/1500, 15*time.Minute, true, make(chan *Request), make(chan *Response), options...)
	// create the temp results channel
	data := make(chan *Data)
	errors := make(chan error)
	// create the request object
	request, _ := NewRquest("GET", fmt.Sprintf("%s/tweets/%s", api.baseURL, id), v, nil)
	// start the requests channel processor
	go queue.processRequests(api)
	// add the 1st request to the channel
	queue.requestsChannel <- request

	// async process the response channel
	go (func(q *Queue, d chan *Data, e chan error, req *Request) {
		// on done close channels
		// close data channel
		defer close(d)
		// close error channel
		defer close(e)

		// listen channel
		for {
			// capture the response and channel state
			res, ok := <-q.responseChannel
			// break the loop if the channel is closed
			if !ok {
				break
			}

			// send the results to the data channel
			d <- &res.Results
			// send errors to error channel
			if res.Error != nil {
				e <- res.Error
			}

			// we are done! break the loop and close the channels
			break
		}
	})(queue, data, errors, request)

	// return the data channel
	return data, errors
}

// GetTweetByID returns a variety of information about a single Tweet specified by the requested ID.
// Endpoint URL: https://api.twitter.com/2/tweets/:id
// Official Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference//get-tweets-id
// Authentication Methods: OAuth 1.0a User Context, OAuth 2.0 Bearer Token
// Rate Limit: 300/15m (app), 900/15m (user)
func (api *Twitter) GetReTweetedBy(id string, v url.Values, options ...QueueOption) (chan *Data, chan error) {
	// create the queue to process requests
	queue := NewQueue(15*time.Minute/1500, 15*time.Minute, true, make(chan *Request), make(chan *Response), options...)
	// create the temp results channel
	data := make(chan *Data)
	errors := make(chan error)
	// create the request object
	request, _ := NewRquest("GET", fmt.Sprintf("%s/tweets/%s/retweeted_by", api.baseURL, id), v, nil)
	// start the requests channel processor
	go queue.processRequests(api)
	// add the 1st request to the channel
	queue.requestsChannel <- request

	// async process the response channel
	go (func(q *Queue, d chan *Data, e chan error, req *Request) {
		// on done close channels
		// close data channel
		defer close(d)
		// close error channel
		defer close(e)

		// listen channel
		for {
			// capture the response and channel state
			res, ok := <-q.responseChannel
			// break the loop if the channel is closed
			if !ok {
				break
			}

			// send the results to the data channel
			d <- &res.Results
			// send errors to error channel
			if res.Error != nil {
				e <- res.Error
			}

			// we are done! break the loop and close the channels
			break
		}
	})(queue, data, errors, request)

	// return the data channel
	return data, errors
}
