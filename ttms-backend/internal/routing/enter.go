/**
 * @Author: lenovo
 * @Description:
 * @File:  enter
 * @Version: 1.0.0
 * @Date: 2023/05/29 8:20
 */

package routing

type group struct {
	User   user
	Email  email
	Movie  movie
	File   file
	Pm     pm
	Tag    tag
	Cinema cinema
	Plan   plan
	Seat   seat
	Ticket ticket
}

var Group = new(group)
