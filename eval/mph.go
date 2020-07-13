//* This file incorporates modified work from https://github.com/alecthomas/mph 
//* covered by the following license: 

	//* Copyright (c) 2014, Alec Thomas
	//* All rights reserved.
	//* 
	//* Redistribution and use in source and binary forms, with or without
	//* modification, are permitted provided that the following conditions are met:
	//* 
	//* - Redistributions of source code must retain the above copyright notice, this
	//* list of conditions and the following disclaimer.
	//* - Redistributions in binary form must reproduce the above copyright notice,
	//* this list of conditions and the following disclaimer in the documentation
	//* and/or other materials provided with the distribution.
	//* - Neither the name of SwapOff.org nor the names of its contributors may
	//* be used to endorse or promote products derived from this software without
	//* specific prior written permission.
	//* 
	//* THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
	//* ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
	//* WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
	//* DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
	//* FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
	//* DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
	//* SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
	//* CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
	//* OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
	//* OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package eval

import (
	"encoding/binary"
	"io"
	"io/ioutil"
)

// chdPoker hash table lookup using the "Compress, Hash, and Displace" algorithm.
// See http://cmph.sourceforge.net/papers/esa09.pdf for details.
type chdPoker struct {
	// Random hash function table.
	r [num_rand_hashes]uint64
	// Array of indices into hash function table r. We assume there aren't
	// more than 2^16 hash functions O_o
	indices [2444]uint16
	// Final table of values.
	//keys   []uint32
	values [4888]uint16
}	

func hasherPoker(data uint32) uint64 {
	var hash uint64 = 14695981039346656037

	for i := 0; i < 4; i++ {
		hash ^= uint64(data & uint32(0xFF<<(8*i)))
		hash *= 1099511628211
	}

	return hash
}
// read a serialized CHDPoker.
func read(r io.Reader) (*chdPoker, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return mmapPoker(b)
}

// Mmap creates a new CHDPoker aliasing the CHDPoker structure over an existing byte region (typically mmapped).
func mmapPoker(b []byte) (*chdPoker, error) {
	c := &chdPoker{}

	// Read vector of hash functions.
	var ndx int = 0

	// binary.LittleEndian.Uint32(b[ndx:ndx + 4])
	// ndx = ndx + 4
	//rl := bi.ReadInt()

	// func (b *sliceReader) ReadUint64Array(n uint64) {
	// 	b.start, b.end = b.end, b.end+n*8
	// 	return unsafeslice.Uint64SliceFromByteSlice(b.b[b.start:b.end])
	// }
	for i := 0; i < num_rand_hashes; i++ {
		c.r[i] = binary.LittleEndian.Uint64(b[ndx + (8 * i): ndx + (8 * i) + 8])
	}
	ndx = ndx + (num_rand_hashes * 8)

	// Read hash function indices.
	// binary.LittleEndian.Uint32(b[ndx:ndx+4])
	// ndx = ndx + 4
	//il := bi.ReadInt()

	for i := 0; i < 2444; i++ {
		c.indices[i] = binary.LittleEndian.Uint16(b[ndx + (2 * i): ndx + (2 * i) + 2])
	}
	ndx = ndx + (2444 * 2)
	

	//el := bi.ReadInt()
	// binary.LittleEndian.Uint32(b[ndx:ndx+4])
	// ndx = ndx + 4


	//c.values = make([]uint16, el)

	for i := 0; i < 4888; i++ {
		c.values[i] = binary.LittleEndian.Uint16(b[ndx + (2 * i): ndx + (2 * i) + 2])
	}


	return c, nil
}

// Get an entry from the hash table.
func (c *chdPoker) get(key uint32) uint16 {

	h := hasherPoker(key) ^ c.r[0]

	ri := c.indices[h % 2444]
	// This can occur if there were unassigned slots in the hash table.
	// if ri >= uint16(len(c.r)) {
	// 	panic(errors.New("TODO: Implement error"))
	// }
	r := c.r[ri]

	ti := (h ^ r) % 4888

	v := c.values[ti]

	return v
}


